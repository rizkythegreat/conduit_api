package route

import (
	"conduit_api/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/users", handler.UserHandlerGetAll)
	r.Get("/users/:id", handler.UserHandlerGetById)
	r.Post("/users", handler.UserHandlerCreate)
	r.Post("/users/bulk", handler.UserHandlerCreateBulk)
}
