package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database []struct {
		User string `mapstructure:"user"`
		Pass string `mapstructure:"password"`
	} `mapstructure:"database"`
}

func main() {
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
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Error unmarshal config: ", err)
	}

	fmt.Println("Port: ", config.Server.Port)

	for _, db := range config.Database {
		fmt.Println("User: ", db.User+" Pass: ", db.Pass+"\n")
	}
}
