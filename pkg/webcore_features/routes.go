package webcore_features

import (
	"github.com/elanticrypt0/gasonline/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(gas *webcore.GasonlineApp) {

	// setup
	setup := gas.Fiber.Group("/api/setup")
	if gas.App.Config.App_setup_enabled {
		setup.Get("/", func(c *fiber.Ctx) error {
			// return webcore_features.Setup(c)
			return Setup(c, gas)
		})
	}

	//status
	status := gas.Fiber.Group("/api/status")
	status.Get("/", func(c *fiber.Ctx) error {
		return Status(c)
	})

	// seeder
	seeder := gas.Fiber.Group("/api/seeder")
	if gas.App.Config.App_setup_enabled {
		seeder.Get("/api/seed/", func(c *fiber.Ctx) error {
			return Seed(c, gas)
		})
		seeder.Get("/api/seed/:table_name", func(c *fiber.Ctx) error {
			return Seed(c, gas)
		})
	}
}
