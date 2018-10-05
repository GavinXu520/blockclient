package main

import (
	"blockclient/src/server/common"
	"blockclient/src/server"
	"blockclient/src/server/timer"
)

func init() {
	common.SetupConfig()

	common.SetErrorDeal()

	//common.SetUpRedis()


	common.InitEthClient()

	timer.SetupTimer()
}

func main() {
	server.SetupServer()
}



