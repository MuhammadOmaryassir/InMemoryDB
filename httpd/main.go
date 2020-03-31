package main

import (
	"inmemorydb/httpd/handler"
	"inmemorydb/platform/inmemorydb"

	"github.com/gin-gonic/gin"
)

func main() {
	memory := inmemorydb.New()

	r := gin.Default()
	r.GET("/memory", handler.Getmemory(memory))
	r.POST("/memory", handler.PostMemory(memory))
	r.Run()

}
