package bank

import (
	"log"
	"strconv"
	db "transaction/internal/adapter/persistence"
	ports "transaction/internal/adapter/port"
	"transaction/internal/core/bank"
)

type BankInterfaceImpl struct {
	userDB ports.UserDB
}

func NewBankInterfaceImpl(db *db.UserDatabase) *BankInterfaceImpl {
	return &BankInterfaceImpl{
		userDB: db,
	}
}

func (b *BankInterfaceImpl) TransferMoney(transfer *bank.MoneyTransfer) error {

	fromUserMoney, err := b.userDB.GetAmountOfUser(transfer.FromAccount)
	if err != nil {
		log.Println("Failed to get amount of user ", err)
		return bank.ErrGettingAmount
	}
	transferAmount, err := strconv.Atoi(transfer.Amount)
	if err != nil {
		log.Println("Failed to convert amount to int ", err)
		return bank.ErrSomethingWentWrong
	}
	if fromUserMoney < int64(transferAmount) {
		log.Println("Insufficient balance")
		return bank.ErrInsufficientBalance
	}

	toUserMoney, err := b.userDB.GetAmountOfUser(transfer.ToAccount)
	if err != nil {
		log.Println("Failed to get amount of user ", err)
		return bank.ErrGettingAmount
	}
	fromUserMoney = fromUserMoney - int64(transferAmount)
	toUserMoney = toUserMoney + int64(transferAmount)

	return nil
}