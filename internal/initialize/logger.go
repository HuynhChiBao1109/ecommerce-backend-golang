package initialize

import (
	"ecommerce-backend-golang/global"
	"ecommerce-backend-golang/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
