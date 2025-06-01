package handler

import (
	"conduit_api/database"
	"conduit_api/model/entity"
	"conduit_api/model/request"
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	db := database.DB

	// Pagination
	skip, _ := strconv.Atoi(ctx.Query("skip", "0"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	// Sorting dengan validasi
	allowedSortColumns := map[string]bool{
		"created_at": true,
		"updated_at": true,
		"name":       true,
		"email":      true,
	}

	requestedSort := ctx.Query("sortBy", "created_at")
	if !allowedSortColumns[requestedSort] {
		requestedSort = "created_at" // Fallback ke default jika invalid
	}
	db = db.Order(requestedSort + " DESC")

	// Keyword search
	keyword := ctx.Query("keyword")
	if keyword != "" {
		db = db.Where("LOWER(name) LIKE ? OR LOWER(email) LIKE ?",
			"%"+strings.ToLower(keyword)+"%",
			"%"+strings.ToLower(keyword)+"%")
	}

	// Filter
	filterParam := ctx.Query("filter")
	if filterParam != "" {
		filters := strings.Split(filterParam, "~")
		for _, f := range filters {
			parts := strings.SplitN(f, ":", 2)
			if len(parts) == 2 {
				field := strings.ToLower(parts[0])
				value := strings.ToLower(parts[1])
				if allowedSortColumns[field] { // Hanya filter field yang diizinkan
					db = db.Where("LOWER("+field+") LIKE ?", "%"+value+"%")
				}
			}
		}
	}

	// Eksekusi query
	if err := db.Offset(skip).Limit(limit).Find(&users).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
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

func UserHandlerCreateBulk(ctx *fiber.Ctx) error {
	var users []entity.User
	if err := ctx.BodyParser(&users); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid request body",
		})
	}

	if len(users) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Empty user list",
		})
	}

	if err := database.DB.Create(&users).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   users,
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
