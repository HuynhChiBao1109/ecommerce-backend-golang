package user

import (
	"ecommerce-backend-golang/internal/controller"
	"ecommerce-backend-golang/internal/repo"
	"ecommerce-backend-golang/internal/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	ur := repo.NewUserRepository()
	us := service.NewUserService(ur)
	userHandlerNonDependency := controller.NewUserController(us)

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.GET("/register", userHandlerNonDependency.Register)
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
