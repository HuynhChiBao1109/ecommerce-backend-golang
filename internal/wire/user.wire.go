package wire

import (
	"ecommerce-backend-golang/internal/controller"
	"ecommerce-backend-golang/internal/repo"
	"ecommerce-backend-golang/internal/service"
)

func InitUserHandler() (*controller.UserController, error) {
	iUserRepository := repo.NewUserRepository()
	iUserService := service.NewUserService(iUserRepository)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
