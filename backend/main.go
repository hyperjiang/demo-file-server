package main

import (
	"log"
	"net/http"
	"path/filepath"

	env "github.com/caarlos0/env/v6"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type config struct {
	StorageDir string `env:"STORAGE_DIR" envDefault:"./files"`
}

func main() {
	cfg := config{}
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

	})

	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// upload the file to specific dst.
		dst := filepath.Join(cfg.StorageDir, file.Filename)
		c.SaveUploadedFile(file, dst)

		c.JSON(http.StatusOK, gin.H{"name": file.Filename, "url": "/files/" + file.Filename})
	})

	router.Run(":9090")
}
