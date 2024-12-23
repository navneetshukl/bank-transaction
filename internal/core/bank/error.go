package bank

import "errors"

var (
	ErrInsufficientBalance error = errors.New("insufficient balance")
	ErrGettingAmount       error = errors.New("sailed to get amount of user")
	ErrSomethingWentWrong  error = errors.New("Something went wrong")
	ErrTransferMoney	   error = errors.New("Failed to transfer money")
)
