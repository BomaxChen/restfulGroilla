package main

import (
	"github.com/BomaxChen/restfulGorilla/api/app"
	"github.com/BomaxChen/restfulGorilla/api/app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
