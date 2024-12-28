package bank

import (
	"fmt"
	"transaction/internal/core/bank"

	"github.com/gofiber/fiber/v2"
)

type Bankhandler struct {
	bankUsecase bank.BankInterface
}

func NewBankHandler(bankUsecase bank.BankInterface) *Bankhandler {
	return &Bankhandler{
		bankUsecase: bankUsecase,
	}
}
func (b *Bankhandler) TransferAmount(c *fiber.Ctx) error {
	var req bank.MoneyTransfer
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong",
			"status":  "failed",
		})
	}
	err = b.bankUsecase.TransferMoney(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "money not transferred",
			"status":  "failed",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Money transferred successfully from %s to %s", req.FromAccount, req.ToAccount),
		"status":  "success",
	})
}
