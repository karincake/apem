package apem

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func (a *app) initConfig() {
	// stated config path will be priority
	cfgPath := "."
	flag.StringVar(&cfgPath, "config-path", ".", "Configuration path (default=./)")

	// main viper config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(cfgPath)
	viper.AddConfigPath("/etc/" + a.CodeName)
	viper.AddConfigPath("$HOME/." + a.CodeName)
	viper.AutomaticEnv()

	// default values
	viper.SetDefault("FullName", "Apem Instance")
	viper.SetDefault("Version", "1.0.0")
	viper.SetDefault("Env", "development")
	viper.SetDefault("LangConf.Active", "en")
	viper.SetDefault("HttpConf.Host", "127.0.0.1")
	viper.SetDefault("HttpConf.Port", "8000")

	// read the file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		panic(err)
	}

	// map to app
	if err := viper.Unmarshal(a); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		panic(err)

	}

	// done
	fmt.Println("Loaded config successfully")
}
