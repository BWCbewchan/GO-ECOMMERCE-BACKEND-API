package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// Grouping routes under /v1/2024
	v1 := r.Group("/v1/2024")
	{
		//
		v1.GET("/ping/:name", Pong)
	}
	return r
}

func Pong(c *gin.Context) {
	// curl http://localhost:8080/v1/2024/ping/bewchan?uid=1234
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong.hhh..ping " + name,
		"uid":     uid,
		"users":   []string{"1", "2", "3"},
	})
}
