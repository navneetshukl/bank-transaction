package ports

type BankDB interface {
	TransferMoney(fromAccount, toAccount string, fromMoney, toMoney int64) error
}
