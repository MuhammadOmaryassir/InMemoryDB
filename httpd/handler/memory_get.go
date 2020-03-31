package handler

import (
	"inmemorydb/platform/inmemorydb"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Getmemory [GET] /memory
func Getmemory(memory *inmemorydb.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		items := memory.GetAll()
		c.JSON(http.StatusOK, items)
	}
}
