package routers

import (
	"ecommerce-backend-golang/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	path := r.Group("/v1")
	{
		path.GET("/ping/:id", controller.NewUserController().GetListuser)
	}

	return r
}
