package webcore_features

import (
	"fmt"

	"github.com/elanticrypt0/gasonline/api/models"
	"github.com/elanticrypt0/gasonline/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func Setup(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	automigrateModels(gas)
	return c.SendString("Setup enabled. Models Migrated.")
}

func SetupOnStartup(gas *webcore.GasonlineApp) {
	fmt.Println("\nDatabase automigration...")
	automigrateModels(gas)
}

func automigrateModels(gas *webcore.GasonlineApp) {
	gas.App.DB.Primary.AutoMigrate(&models.Category{})
}
