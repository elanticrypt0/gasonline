package main

import (
	"log"

	"github.com/elanticrypt0/gasonline/api/routes"
	"github.com/elanticrypt0/gasonline/pkg/webcore"
	"github.com/elanticrypt0/gasonline/pkg/webcore_features"
	"github.com/elanticrypt0/go4it"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app_config := go4it.NewApp("./appconfig")

	// make the connection
	app_config.Connect2Db("local")
	app_config.DB.SetPrimaryDB(0)

	gas := webcore.GasonlineApp{
		App: &app_config,
		Fiber: fiber.New(fiber.Config{
			Prefork:       false,
			CaseSensitive: true,
			StrictRouting: true,
			ServerHeader:  "Fiber",
			AppName:       app_config.Config.App_name,
			// access
			DisableStartupMessage: false,
			PassLocalsToViews:     true,
		}),
	}

	webcore.MiddlewareSetup(&gas)

	gas.PrintAppInfo()

	// Routes setup

	// features routes
	routes.SetupApiRoutes(&gas)

	// webcore setup routes
	if gas.App.Config.App_setup_enabled {
		// webcore features setup
		webcore_features.SetupRoutes(&gas)

		webcore_features.SetupOnStartup(&gas)
	}
	// static routes
	webcore.SetupStaticRoutes(gas.Fiber)

	portAsStr := gas.GetPortAsStr()

	// go4it.OpenInBrowser("http://" + gas.GetAppUrl())

	log.Fatal(gas.Fiber.Listen(gas.GetAppUrl()), "Server is running on port "+portAsStr)

}
