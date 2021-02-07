package util

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

// LoadConfig 读取配置文件
func LoadConfig() (*viper.Viper, error) {
	config := viper.New()
	config.AddConfigPath(".")
	config.AddConfigPath("~/Azeroth/Hellscream/config")
	config.AddConfigPath("/Azeroth/Hellscream/config")
	//for run test
	config.AddConfigPath("../../")
	config.SetConfigName("hellscream_conf")
	config.SetConfigType("yaml")

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	return config, nil
}
