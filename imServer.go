package main

import (
	config "github.com/config"
	controller "github.com/controller"
)

func main() {
	appConfig := config.AppConfig{}
	appConfig.Init()
	controller.ServerSetup(appConfig, "9000")
}
