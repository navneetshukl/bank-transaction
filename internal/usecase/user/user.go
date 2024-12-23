package user

import (
	"log"
	db "transaction/internal/adapter/persistence"
	ports "transaction/internal/adapter/port"
	"transaction/internal/core/user"
)

type UserInterfaceImpl struct {
	db ports.UserDB
}

func NewUserInterfaceImpl(db *db.UserDatabase) *UserInterfaceImpl {
	return &UserInterfaceImpl{
		db: db,
	}
}

func (u *UserInterfaceImpl) CreateAccount(userDet *user.User) (string, error) {

	if len(userDet.Email) == 0 || len(userDet.Name) == 0 || len(userDet.Phone) == 0 || len(userDet.Bank) == 0 {
		log.Println("Invalid User detail")
		return "", user.ErrInvalidUser
	}

	// generate account number of 10 digits for user
	accountNumber, err := user.GenerateUniqueRandomValue(10)
	if err != nil {
		log.Println("Failed to generate account number")
		return "", user.ErrGeneratingAccountNumber
	}
	userDet.Account = accountNumber
	userDet.Money = 0
	err = u.db.InsertUser(userDet)
	if err != nil {
		log.Println("Failed to insert user")
		return "", user.ErrCreatingUser
	}
	return accountNumber, nil
}

func (u *UserInterfaceImpl) UpdateAmount(account string, money int64) error {

	// Get previous amount from DB

	amount, err := u.db.GetAmountOfUser(account)
	if err != db.ErrNoUserFound {
		log.Println("Failed to get amount of user")
		return user.ErrGettingAmount
	}
	amount = amount + money
	err = u.db.UpdateAmount(account, amount)
	if err != nil {
		log.Println("Failed to update amount of user")
		return user.ErrUpdatingAmount
	}

	return nil
}
