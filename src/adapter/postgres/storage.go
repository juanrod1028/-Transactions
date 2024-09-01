package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/juanrod1028/Transactions/src/usecase/models"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(host string, port int, user string, password string, dbname string) (*PostgresStore, error) {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}
func (s *PostgresStore) Init() error {
	return s.createTables()
}
func (s *PostgresStore) createTables() error {
	userTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		identification VARCHAR(50) PRIMARY KEY,
		email VARCHAR(100),
		created_at TIMESTAMPTZ
	);`

	transactionTableQuery := `
	CREATE TABLE IF NOT EXISTS transactions (
		id SERIAL PRIMARY KEY,
		transaction_id VARCHAR(50),
		user_id VARCHAR(50) REFERENCES users(identification),
		date VARCHAR(10),
		transaction FLOAT,
		created_at TIMESTAMPTZ
	);`

	_, err := s.db.Exec(userTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	_, err = s.db.Exec(transactionTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create transactions table: %v", err)
	}
	return err
}
func (s *PostgresStore) CreateUser(user *models.User) error {
	var existingID string
	query := `SELECT identification FROM users WHERE identification = $1`
	err := s.db.QueryRow(query, user.Identification).Scan(&existingID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check if user exists: %v", err)
	}

	if existingID != "" {
		log.Printf("User with ID %s already exists. Skipping creation.", user.Identification)
		return nil
	}

	insertQuery := `INSERT INTO users 
		(identification, email, created_at)
		VALUES 
		($1, $2, $3)`
	_, err = s.db.Exec(insertQuery, user.Identification, user.Email, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}
func (s *PostgresStore) CreateTransactions(user *models.User) error {
	query := `INSERT INTO transactions 
		(transaction_id, user_id, date, transaction, created_at)
		VALUES 
		($1, $2, $3, $4, $5)`

	for _, tx := range user.Transactions {
		fmt.Print(tx)
		_, err := s.db.Exec(query, tx.Id, user.Identification, tx.Date, tx.Transactions, tx.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to insert transaction %d: %v", tx.Id, err)
		}
	}

	return nil
}
func (s *PostgresStore) GetTransactionById(id int) (*models.User, error) {
	rows, err := s.db.Query("SELECT transaction_id, date, transaction, created_at FROM transactions WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transactions

	for rows.Next() {
		var tx models.Transactions

		err := rows.Scan(&tx.Id, &tx.Date, &tx.Transactions, &tx.CreatedAt)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, tx)
	}

	if len(transactions) == 0 {
		return nil, fmt.Errorf("transaction for user %d not found", id)
	}

	var user models.User
	userQuery := `SELECT identification, email FROM users WHERE identification = $1`
	err = s.db.QueryRow(userQuery, id).Scan(&user.Identification, &user.Email)
	if err != nil {
		return nil, err
	}

	user.Transactions = transactions

	return &user, nil
}
