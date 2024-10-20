package service

import (
	"github.com/felipeazsantos/personal-finance-go/internal/models"
	"github.com/felipeazsantos/personal-finance-go/internal/repository"
)

type FinancialTransactionService struct {
	repo repository.Repository
}

func (s *FinancialTransactionService) GetFinancialTransactionByID(id uint64) (*models.FinancialTransaction, error) {
	return s.repo.GetFinancialTransactionByID(id)
}

func (s *FinancialTransactionService) InsertFinancialTransaction(ft *models.FinancialTransaction) error {
	return s.repo.InsertFinancialTransaction(ft)
}

func (s *FinancialTransactionService) UpdateFinancialTransaction(ft *models.FinancialTransaction) error {
	return s.repo.UpdateFinancialTransaction(ft)
}

func (s *FinancialTransactionService) DeleteFinancialTransactionByID(id uint64) error {
	return s.repo.DeleteFinancialTransactionByID(id)
}

func (s *FinancialTransactionService) GetAllFinancialTransaction(limit, page uint64) ([]models.FinancialTransaction, error) {
	return s.repo.GetAllFinancialTransaction(limit, page)
}
