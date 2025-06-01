package handler

import (
	"conduit_api/database"
	"conduit_api/model/entity"
	"conduit_api/model/request"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	if len(users) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "No users found",
		})
	}
	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   users,
		"count":  len(users),
	})
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	var user request.UserCreateRequest
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}

	newUser := entity.User{
		Name:  user.Name,
		Email: user.Email,
	}
	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":   "error",
			"messsage": errCreateUser,
		})
	}
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Success create user",
		"data":    user,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid ID parameter",
		})
	}

	var user entity.User
	err = database.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "User not found",
		})
	} else if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}
