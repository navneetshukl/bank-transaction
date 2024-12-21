package user

import (
	"fmt"
	"transaction/internal/core/user"
)

type UserInterfaceImpl struct{}

func NewUserInterfaceImpl() *UserInterfaceImpl {
	return &UserInterfaceImpl{}
}

func (u *UserInterfaceImpl) CreateAccount(userDet *user.User) error {

	if len(userDet.Email) == 0 || len(userDet.Name) == 0 || len(userDet.Phone) == 0 || len(userDet.Bank) == 0 {
		return user.ErrInvalidUser
	}

	// generate account number of 10 digits for user
	accountNumber,err:=user.GenerateUniqueRandomValue(10)
	if err!=nil{
		return user.ErrGeneratingAccountNumber
	}

	fmt.Println("Account Number:", accountNumber)


	return nil
}
