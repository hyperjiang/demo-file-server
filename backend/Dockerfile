FROM golang:1.13-alpine as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /go/src/backend

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /tmp/main .

FROM gcr.io/distroless/base

COPY --from=builder /tmp/main /main

ENTRYPOINT ["/main"]
