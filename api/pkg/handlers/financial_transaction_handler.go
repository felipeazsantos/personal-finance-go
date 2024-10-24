package handlers

import (
	"net/http"
	"strconv"

	"github.com/felipeazsantos/personal-finance-go/pkg/utils"
)

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

	h.FinancialTransactionService.GetFinancialTransactionByID(idUint64)
}
