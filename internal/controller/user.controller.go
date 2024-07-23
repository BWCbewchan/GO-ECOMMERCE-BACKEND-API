package controller

import (
	"net/http"

	"github.com/BWCbewchan/GO-ECOMMERCE-BACKEND-API/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserRepo(),
	}
}

//uc user controller
//uc user service

//controller -> service -> repo -> model -> dbs
func (uc *UserController) GetUserByID(c *gin.Context) {
	// curl http://localhost:8080/v1/2024/ping/bewchan?uid=1234
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInforUser(),
		"users":   []string{"1", "2", "3"},
	})
}
