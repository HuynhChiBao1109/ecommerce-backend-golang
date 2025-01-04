package initialize

import (
	"ecommerce-backend-golang/global"
	"fmt"

	"go.uber.org/zap"
)

func Run() {
	// Initialize the configuration
	InitConfig()
	fmt.Println("Config initialized", global.Config.MySQL.Host)
	InitLogger()
	global.Logger.Info("Logger initialized", zap.String("log level", "sucess"))
	InitMySQL()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
