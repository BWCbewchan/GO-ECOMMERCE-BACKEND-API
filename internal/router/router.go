package router

import (
	"fmt"

	c "github.com/BWCbewchan/GO-ECOMMERCE-BACKEND-API/internal/controller"
	"github.com/BWCbewchan/GO-ECOMMERCE-BACKEND-API/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After  --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After  --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After  --> CC")
}
func NewRouter() *gin.Engine {
	r := gin.Default()
	//use the middleware
	r.Use(middlewares.AuthenMiddleware(), BB(), CC)
	// Grouping routes under /v1/2024
	v1 := r.Group("/v1/2024")
	{
		//
		v1.GET("/ping/:name", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserByID)
	}
	return r
}
