package db

import (
	"database/sql"
	"transaction/internal/core/user"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{DB: db}
}

func (u *Database) InsertUser(userDet *user.User) error {
	defer u.DB.Close()

	query := `Insert into users(email,name,phone,bank,account_number) values($1,$2,$3,$4,$5)`
	_, err := u.DB.Exec(query, userDet.Email, userDet.Name, userDet.Phone, userDet.Bank, userDet.Account)
	if err != nil {
		return err
	}
	return nil
}

func (u *Database) UpdateAmount(account string, money int64) error {
	defer u.DB.Close()

	query := `UPDATE users SET amount=$1 WHERE account_number=$2`
	_, err := u.DB.Exec(query, money, account)
	if err == sql.ErrNoRows {
		return ErrNoUserFound
	}
	if err != nil {
		return err
	}

	return nil
}

func (u *Database) GetUserCount(account string) (int, error) {
	defer u.DB.Close()
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

func (u *Database) GetAmountOfUser(account string) (int64, error) {
	query := `SELECT amount from users where account_number=$1`
	var amount int64
	err := u.DB.QueryRow(query, account).Scan(&amount)
	if err == sql.ErrNoRows {
		return 0, ErrNoUserFound
	}
	if err != nil {
		return 0, err
	}
	return amount, nil

}
