package routes

import (
	"transaction/internal/interface/handler/bank"
	"transaction/internal/interface/handler/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(userhandler *user.Userhandler,bankhandler *bank.Bankhandler) *fiber.App {
	router := fiber.New()

	router.Post("/api/create", userhandler.CreateUser)
	router.Post("/api/update-amount", userhandler.UpdateAmount)
	router.Post("/api/transfer",bankhandler.TransferAmount)

	return router
}
