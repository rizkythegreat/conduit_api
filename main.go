package main

import (
	"conduit_api/database"
	"conduit_api/database/migration"
	"conduit_api/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initial Database
	database.DatabaseInit()
	migration.RunMigration()

	// Initial Fiber
	app := fiber.New()

	// Initial Route
	route.RouteInit(app)

	app.Listen(":3000")
}
