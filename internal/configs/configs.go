package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfigsViper() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	// set default
	// viper.SetDefault("web.port", "80")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
		os.Exit(1)
	}
}
