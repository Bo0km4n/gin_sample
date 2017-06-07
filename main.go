package main

import (
	"flag"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"router"
)

var env = flag.String("e", "dev", "環境変数")

func init() {
	viper.AddConfigPath("./config")
	switch *env {
	case "dev":
		viper.SetConfigName("dev")
	default:
		viper.SetConfigName("dev")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config error", err)
	}
	
}

func main() {
	api := gin.New()
	api.Use(gin.Recovery())

	router.V1(api)
	router.HealthCheck(api)

	server := &http.Server{
		Addr:    viper.GetString("port"),
		Handler: api,
	}

	gracehttp.Serve(server)
}
