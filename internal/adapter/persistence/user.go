package db

import (
	"database/sql"
	"log"
	"transaction/internal/core/user"
)

type UserDatabase struct {
	DB *sql.DB
}

func NewUserDatabase(db *sql.DB) *UserDatabase {
	return &UserDatabase{DB: db}
}

func (u *UserDatabase) InsertUser(userDet *user.User) error {
	//defer u.DB.Close()

	query := `Insert into users(email,name,phone,bank,account_number,money) values($1,$2,$3,$4,$5,$6)`
	_, err := u.DB.Exec(query, userDet.Email, userDet.Name, userDet.Phone, userDet.Bank, userDet.Account, userDet.Money)
	if err != nil {
		log.Println("Error inserting user", err)
		return err
	}
	return nil
}

func (u *UserDatabase) UpdateAmount(account string, money int64) error {
	//defer u.DB.Close()

	query := `UPDATE users SET money=$1 WHERE account_number=$2`
	_, err := u.DB.Exec(query, money, account)
	if err == sql.ErrNoRows {
		return ErrNoUserFound
	}
	if err != nil {
		return err
	}

	return nil
}

func (u *UserDatabase) GetUserCount(account string) (int, error) {
	//defer u.DB.Close()
	query := `SELECT count(account_number) from users where account_number=$1`
	var count int
	err := u.DB.QueryRow(query, account).Scan(&count)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *UserDatabase) GetAmountOfUser(account string) (int64, error) {
	query := `SELECT money from users where account_number=$1`
	var money int64
	err := u.DB.QueryRow(query, account).Scan(&money)
	if err == sql.ErrNoRows {
		return 0, ErrNoUserFound
	}
	if err != nil {
		return 0, err
	}
	return money, nil

}
