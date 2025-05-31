package handler

import (
	"conduit_api/database"
	"conduit_api/model/entity"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	err := database.DB.Find(&users).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}
	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   users,
		"count":  len(users),
	})
}
