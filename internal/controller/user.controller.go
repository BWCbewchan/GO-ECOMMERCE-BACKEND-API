package controller

import (
	"github.com/BWCbewchan/GO-ECOMMERCE-BACKEND-API/internal/service"
	"github.com/BWCbewchan/GO-ECOMMERCE-BACKEND-API/pkg/response"
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

// controller -> service -> repo -> model -> dbs
func (uc *UserController) GetUserByID(c *gin.Context) {
	// if err != nil {
	// 	return response.ErrorResponse(c, 20003, "no need!!")

	// }
	// return response.SuccessResponse(c, 20001, []string{"tipjs", "m10", "anonystick"})
	// curl http://localhost:8080/v1/2024/ping/bewchan?uid=1234
	response.SuccessResponse(c, 20001, []string{"tipjs", "m10", "anonystick"})
	// response.ErrorResponse(c, 20003, "no need!!")
}
