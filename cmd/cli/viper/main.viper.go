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
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		// DbName string `mapstructure:"dbName"`
	}
}

func main() {
	v := viper.New()
	v.AddConfigPath("./config/") // Path to config file
	v.SetConfigName("local")     // Name of config file (without extension)
	v.SetConfigType("yaml")      // Config file type

	// Read config
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// Read server config
	fmt.Printf("server port:: %d \n", v.GetInt("server.port"))
	fmt.Printf("server jwtkey:: %s \n", v.GetString("security.jwt.key"))

	//configure structure
	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	// fmt.Printf("database user:: %s \n", config.Database[0].User)
	fmt.Println("Config port:: ", config.Server.Port)

	for _, db := range config.Database {
		fmt.Printf("database user: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)
	}
}
