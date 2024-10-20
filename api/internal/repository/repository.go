package repository

import (
	"database/sql"

	"github.com/felipeazsantos/personal-finance-go/internal/models"
)

type Repository interface {
	GetFinancialTransactionByID(ID uint64) (*models.FinancialTransaction, error)
	InsertFinancialTransaction(ft *models.FinancialTransaction) error
	GetAllFinancialTransaction(limit, page uint64) (result []models.FinancialTransaction, err error)
	DeleteFinancialTransactionByID(id uint64) error
	UpdateFinancialTransaction(ftUpdated *models.FinancialTransaction) error
}

type Postgres struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &Postgres{
		db: db,
	}
}
