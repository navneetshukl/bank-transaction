package ports

import "transaction/internal/core/user"

type UserDB interface {
	InsertUser(user *user.User) error
	UpdateAmount(account string, money int64) error
	GetUserCount(account string)(int,error)
	GetAmountOfUser(account string)(int64,error)
}
