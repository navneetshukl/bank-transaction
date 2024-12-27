package bank

import "transaction/internal/core/bank"

type Bankhandler struct {
	bankUsecase bank.BankInterface
}

func NewBankHandler(bankUsecase bank.BankInterface) *Bankhandler {
	return &Bankhandler{
		bankUsecase: bankUsecase,
	}
}