package db

import "database/sql"

type BankDatabase struct {
	DB *sql.DB
}

func NewBankDatabase(db *sql.DB) *BankDatabase {
	return &BankDatabase{DB: db}
}

func (b *BankDatabase) TransferMoney(fromAccount, toAccount string, fromMoney, toMoney int64) error {
	defer b.DB.Close()
	tx, err := b.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := `UPDATE users SET amount=$1 WHERE account_number=$2`
	_, err = tx.Exec(query, fromMoney, fromAccount)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `UPDATE users SET amount=$1 WHERE account_number=$2`
	_, err = tx.Exec(query, toMoney, toAccount)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
