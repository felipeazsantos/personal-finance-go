package handlers

import (
	"net/http"
	"strconv"

	"github.com/felipeazsantos/personal-finance-go/internal/service"
	"github.com/felipeazsantos/personal-finance-go/pkg/utils"
)

type Handlers struct {
	Service *service.Service
}

var financialTransactionService *service.FinancialTransactionService

func NewHandlers(svce *service.Service) *Handlers {
	financialTransactionService = svce.FinancialTransaction
	return &Handlers{
		Service: svce,
	}
}

func (h *Handlers) GetFinancialTransactionByID(w http.ResponseWriter, r *http.Request) {
	id := utils.ExtractRequestParam(r, "id")

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	financialTransactionService.GetFinancialTransactionByID(idUint64)
}
