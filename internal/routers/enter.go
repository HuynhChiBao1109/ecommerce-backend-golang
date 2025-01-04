package routers

import (
	"ecommerce-backend-golang/internal/routers/manage"
	"ecommerce-backend-golang/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
