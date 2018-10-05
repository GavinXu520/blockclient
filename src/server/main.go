package server

import (
	"github.com/spf13/viper"
	"gopkg.in/kataras/iris.v5"
	"blockclient/src/server/route/api"
)

func SetupServer() {

	port := viper.GetString("server.port")
	host := viper.GetString("server.host")
	api.Api()
	iris.Listen(host + ":" + port)

}
