package ports

import (
	"github.com/juanrod1028/Transactions/src/usecase/models"
)

type Storage interface {
	CreateUser(user *models.User) error
	CreateTransactions(user *models.User) error
	GetTransactionById(id int) (*models.User, error)
}
