package main

import (
	"database/sql"
	"fmt"
	"github/j1mb0b/gobank/config"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(id int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(cfg *config.Config) (*PostgresStore, error) {
	db, err := sql.Open("postgres", cfg.PG.URL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id serial NOT NULL PRIMARY KEY,
		firstname VARCHAR(50) NOT NULL,
		lastname VARCHAR(50) NOT NULL,
		accountNumber serial,
		accountBalance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgresStore) CreateAccount(account *Account) error {

	// Insert the account into the database.
	query, err := s.db.Query(`INSERT INTO account
	(firstname, lastname, accountNumber, accountBalance, created_at)
	VALUES ($1, $2, $3, $4, $5)`,
		account.FirstName,
		account.LastName,
		account.AccNumber,
		account.AccBalance,
		account.CreatedAt)
	if err != nil {
		return err
	}
	query.Close()

	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Exec(`DELETE FROM account WHERE ID = $1`)
	return err
}

func (s *PostgresStore) UpdateAccount(account *Account) error {
	return nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {

	// Insert the account into the database.
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	row, err := s.db.Query("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		return scanIntoAccount(row)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := Account{}
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.AccNumber,
		&account.AccBalance,
		&account.CreatedAt)

	return &account, err
}
