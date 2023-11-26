package routes

import (
	"github.com/elanticrypt0/gasonline/pkg/webcore"
)

func SetupFeaturesRoutes(gas *webcore.GasonlineApp) {
	api := gas.Fiber.Group("/api")
	// categories
	categoriesRoutes(gas, api)
}
