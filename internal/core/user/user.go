package user

type UserInterface interface {
	CreateAccount(user *User) (string,error)
	UpdateAmount(account string, money int64) error
}
