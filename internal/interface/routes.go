package routes

import (
	"transaction/internal/interface/handler/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(userhandler *user.Userhandler) *fiber.App {
	router := fiber.New()

	router.Post("/api/create", userhandler.CreateUser)
	router.Post("/api/update-amount", userhandler.UpdateAmount)

	return router
}
