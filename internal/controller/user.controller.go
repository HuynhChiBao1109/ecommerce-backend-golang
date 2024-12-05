package controller

import (
	"ecommerce-backend-golang/internal/service"
	"ecommerce-backend-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// Controller => Service => Repo => Model => DB
func (uc *UserController) GetListuser(c *gin.Context) {
	response.SuccessReponse(c, 20001, uc.userService.GetListUser())
}
