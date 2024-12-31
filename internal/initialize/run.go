package initialize

import (
	"ecommerce-backend-golang/global"
	"fmt"
)

func Run() {
	// Initialize the configuration
	InitConfig()
	fmt.Println("Config initialized", global.Config.MySQL.Host)
	InitLogger()
	InitMySQL()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
