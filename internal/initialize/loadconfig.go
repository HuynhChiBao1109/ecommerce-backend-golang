package initialize

import (
	"ecommerce-backend-golang/global"
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	// Initialize viper
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("development")
	viper.SetConfigType("yaml")

	// // Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	} else {
		fmt.Println("Config file loaded successfully")
	}

	// Get the value from the configuration file
	// fmt.Println("Name: ", viper.GetString("database.user"))
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Error unmarshal config: ", err)
	}
}
