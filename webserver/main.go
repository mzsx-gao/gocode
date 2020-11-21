package main

import (
	"gocode/webserver/service"
	"gocode/webserver/web"
	"gocode/webserver/web/controller"
)

const (
	configFile  = "config.yaml"
	initialized = false
	SimpleCC    = "simplecc"
)

func main() {
	serviceSetup := service.ServiceSetup{
		ChaincodeID: SimpleCC,
		Client:      nil,
	}

	//启动web服务器
	app := controller.Application{
		Fabric: &serviceSetup,
	}
	web.WebStart(&app)

}
