package models

type FinancialTransaction struct {
	Base
	Type       string  `json:"type"`
	Amount     float64 `json:"amount"`
	CategoryID uint64  `json:"category_id"`
}
