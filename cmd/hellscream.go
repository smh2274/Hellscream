package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smh2274/Hellscream/internal/util"
	"log"
	"net/http"
)

func main() {
	// 捕获panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	// 读取配置文件
	viper, err := util.LoadConfig()
	if err != nil {
		log.Fatalf("load config fail: %v", err)
	}

	gin.SetMode(viper.GetString("mode"))

	engine := gin.New()

	// 注册开放访问文件
	pubMap := viper.GetStringMapString("file.public")
	for key, val := range pubMap {
		engine.StaticFS(fmt.Sprintf("/hellscream/public/%s", key), http.Dir(val))
	}

	// 注册需要token才能访问文件
	protectMap := viper.GetStringMapString("file.protect")
	for key, val := range protectMap {
		engine.StaticFS(fmt.Sprintf("/hellscream/protect/%s", key), http.Dir(val))
	}

	address := viper.GetString("server.address")
	port := viper.GetString("server.port")

	certFile := viper.GetString("ssl.cert")
	keyFile := viper.GetString("ssl.key")

	if err := engine.RunTLS(fmt.Sprintf("%s:%s", address, port), certFile, keyFile); err != nil {
		log.Fatalf("start hellscream server fail: %v", err)
	}
	log.Printf("start hellscream server success: %s:%s", address, port)
}
