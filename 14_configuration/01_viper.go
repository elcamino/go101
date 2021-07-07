package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

func Viper() {
	viper.SetDefault("ContentDir", "/content")

	viper.SetEnvPrefix("wave")
	viper.AutomaticEnv()

	viper.SetConfigName("config")       // name of config file (without extension)
	viper.SetConfigType("yaml")         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/myapp/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.myapp") // call multiple times to add many search paths
	viper.AddConfigPath("./data/")      // optionally look for config in the working directory
	err := viper.ReadInConfig()         // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	fmt.Printf("ContentDir: %s\n", viper.GetString("ContentDir"))
	fmt.Printf("Host: %v\n", viper.Get("host"))
	fmt.Printf("Loglevel: %s\n", viper.Get("loglevel"))
	fmt.Printf("Host.Address: %s\n", viper.GetString("host.address"))
}
