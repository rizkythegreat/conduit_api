package handler

import (
	"conduit_api/database"
	"conduit_api/model/entity"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ArticleHandlerGetAll(ctx *fiber.Ctx) error {
	var articles []entity.Article
	err := database.DB.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "Articles not found",
		})
	} else if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   articles,
		"count":  len(articles),
	})
}
