package user

import (
	"net/http"
	"strconv"
	"transaction/internal/core/user"

	"github.com/gofiber/fiber/v2"
)

type Userhandler struct {
	userUsecase user.UserInterface
}

func NewUserHandler(userUsecase user.UserInterface) *Userhandler {
	return &Userhandler{
		userUsecase: userUsecase,
	}
}

func (u *Userhandler) CreateUser(c *fiber.Ctx) error {
	var user user.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong",
			"status":  "failed",
			"error":err.Error(),
		})
	}

	accNum, err := u.userUsecase.CreateAccount(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create account",
			"status":  "failed",
			"error":err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "account created successfully",
		"status":  "success",
		"data": fiber.Map{
			"account_number": accNum,
		},
	})

}

func (u *Userhandler) UpdateAmount(c *fiber.Ctx) error {
	type Req struct {
		AccountNumber string `json:"account_number"`
		Amount        string `json:"amount"`
	}

	var userReq Req

	err := c.BodyParser(&userReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid request",
			"status":  "failed",
			"error":err.Error(),
		})
	}

	// convert amount from string to int

	amount, err := strconv.Atoi(userReq.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong",
			"status":  "failed",
			"error":err.Error(),
		})
	}

	err = u.userUsecase.UpdateAmount(userReq.AccountNumber, int64(amount))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update amount",
			"status":  "failed",
			"error":err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "amount updated successfully",
		"status":  "success",
		"data":    fiber.Map{},
	})
}
