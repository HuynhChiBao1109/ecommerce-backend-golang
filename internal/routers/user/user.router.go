package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.GET("/register")
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
