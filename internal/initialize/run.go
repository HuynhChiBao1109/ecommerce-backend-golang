package initialize

import (
	"ecommerce-backend-golang/global"

	"go.uber.org/zap"
)

func Run() {
	// Initialize the configuration
	InitConfig()
	InitLogger()
	global.Logger.Info("Logger initialized", zap.String("log level", "sucess"))
	InitMySQL()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
