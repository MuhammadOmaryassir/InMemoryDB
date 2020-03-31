package handler

import (
	"inmemorydb/platform/inmemorydb"
	"net/http"

	"github.com/gin-gonic/gin"
)

// memoryPostBody
type memoryPostBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// PostMemory add a new item to memory
func PostMemory(memory *inmemorydb.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := memoryPostBody{}
		c.Bind(&body)
		item := inmemorydb.Item{
			Key:   body.Key,
			Value: body.Value,
		}
		memory.Add(item)
		c.JSON(http.StatusNoContent, item)
	}
}
