package bank

type MoneyTransfer struct{
	FromAccount string `json:"from_account"`
	ToAccount string `json:"to_account"`
	Amount string `json:"amount"`
}

type BankInterface interface {
	TransferMoney(transfer *MoneyTransfer) error
}