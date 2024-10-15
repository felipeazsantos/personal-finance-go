package repository

import (
	"context"
	"time"

	"github.com/felipeazsantos/personal-finance-go/internal/models"
)

// InsertFinancialTransaction insert a new financial transaction to database
func (p *Postgres) InsertFinancialTransaction(ft *models.FinancialTransaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := p.db.ExecContext(ctx,
		`insert into financial_trasaction (type, category_id, amount, created_at, updated_at)
			values (?, ?, ?, ?, ?)`,
		ft.Type,
		ft.CategoryID,
		ft.Amount,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

// GetFinancialTransactionByID returns a financial transaction by id
func (p *Postgres) GetFinancialTransactionByID(id uint64) (*models.FinancialTransaction, error) {
	ft := &models.FinancialTransaction{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select id, type, amount, category_id, created_at, updated_at
				from financial_transaction
				where id = ?
				limit 1`

	row := p.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&ft.ID,
		&ft.Type,
		&ft.Amount,
		&ft.CategoryID,
		&ft.CreatedAt,
		&ft.UpdatedAt,
	)

	if err != nil {
		return ft, err
	}

	return ft, nil
}
