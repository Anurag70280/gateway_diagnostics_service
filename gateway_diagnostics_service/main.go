package main

import (
	"golang_service_template/utils"
	//please leave this space so that utils package gets imported first. This is needed so that env variables get loaded first!
	"golang_service_template/app"
)

func main() {

	utils.ShowServiceInfo()

	app.StartApp()
}
