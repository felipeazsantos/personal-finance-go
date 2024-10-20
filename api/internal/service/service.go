package service

import (
	"github.com/felipeazsantos/personal-finance-go/internal/repository"
)

type Service struct {
	FinancialTransaction *FinancialTransactionService
}

func New(r repository.Repository) *Service {
	return &Service{
		FinancialTransaction: &FinancialTransactionService{repo: r},
	}
}
