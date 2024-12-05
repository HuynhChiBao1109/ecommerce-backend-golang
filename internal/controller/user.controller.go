package controller

import (
	"ecommerce-backend-golang/internal/service"
	"net/http"

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
	// Get id
	id := c.Param("id")
	test := c.Query("test")
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetListUser(),
		"status":  http.StatusOK,
		"id":      id,
		"param":   test,
	})
}
