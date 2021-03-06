package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	env "github.com/caarlos0/env/v6"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type config struct {
	StorageDir string `env:"STORAGE_DIR" envDefault:"./files"`
	Port       int    `env:"PORT" envDefault:"9090"`
}

type fileInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var cfg config

func newFileInfo(filename string) fileInfo {
	return fileInfo{
		Name: filename,
		URL:  fmt.Sprintf("http://localhost:%d/files/%s", cfg.Port, filename),
	}
}

func main() {
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	// set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB

	// cors middleware
	router.Use(cors.Default())

	// serving static files
	router.Static("/files", cfg.StorageDir)

	// get file list
	router.GET("/list", func(c *gin.Context) {
		var files []fileInfo
		filepath.Walk(cfg.StorageDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Printf("fail to access path %q: %v\n", path, err)
				return err
			}
			if !info.IsDir() && !strings.HasPrefix(info.Name(), ".") {
				files = append(files, newFileInfo(info.Name()))
			}
			return nil
		})

		c.JSON(http.StatusOK, files)
	})

	// upload a file
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")

		// upload the file to specific dst.
		dst := filepath.Join(cfg.StorageDir, file.Filename)
		c.SaveUploadedFile(file, dst)

		c.JSON(http.StatusOK, newFileInfo(file.Filename))
	})

	// delete a file
	router.DELETE("/files/:path", func(c *gin.Context) {
		path := c.Param("path")
		dst := filepath.Join(cfg.StorageDir, path)
		err := os.Remove(dst)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
