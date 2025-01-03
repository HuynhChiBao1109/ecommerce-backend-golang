package initialize

import (
	"ecommerce-backend-golang/internal/controller"
	"ecommerce-backend-golang/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthMiddleware())

	path := r.Group("/v1")
	{
		path.GET("/ping/:id", controller.NewUserController().GetListuser)
	}

	return r
}
