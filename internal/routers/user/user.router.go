package user

import (
	"ecommerce-backend-golang/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitRouterUserHandler()

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.GET("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}

	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(LimitMiddleware())
	// userRouterPrivate.Use(AuthMiddleware())
	// userRouterPrivate.Use(PermissionMiddleware())
	{
		userRouterPrivate.GET("/get_info")
	}
}
