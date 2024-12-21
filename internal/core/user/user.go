package user

type UserInterface interface {
	CreateAccount(user *User) error
	UpdateAccount(account string, money int64) error
}
