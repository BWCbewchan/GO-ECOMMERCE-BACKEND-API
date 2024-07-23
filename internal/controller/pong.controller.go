package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (uc *PongController) Pong(c *gin.Context) {
	// curl http://localhost:8080/v1/2024/ping/bewchan?uid=1234
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong.hhh..ping " + name,
		"uid":     uid,
		"users":   []string{"1", "2", "3"},
	})
}
