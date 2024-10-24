package handlers

import (
	"github.com/felipeazsantos/personal-finance-go/internal/service"
)

type Handlers struct {
	FinancialTransactionService *service.FinancialTransactionService
}

func NewHandlers(svce *service.Service) *Handlers {
	return &Handlers{
		FinancialTransactionService: svce.FinancialTransaction,
	}
}
