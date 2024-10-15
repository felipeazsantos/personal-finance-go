package repository

import (
	"database/sql"

	"github.com/felipeazsantos/personal-finance-go/internal/models"
)

type Repository interface {
	GetFinancialTransactionByID(ID uint64) (*models.FinancialTransaction, error)
	InsertFinancialTransaction(ft *models.FinancialTransaction) error
}

type Postgres struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &Postgres{
		db: db,
	}
}
