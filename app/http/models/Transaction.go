package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserId             uint      `json:"user_id"`
	AccountId          uint      `json:"account_id"`
	CategoryId         uint      `json:"catgory)id"`
	Amount             float64   `json:"amount"`
	Description        string    `json:"description"`
	TransactionDate    time.Time `json:"transaction_date"`
	TransactionType    string    `json:"transaction_type"`
	ReferenceNumber    string    `json:"reference_number"`
	RecurringFrequency string    `json:"recurring_frequency"`
	IsOpeningBalance   bool      `json:"is_opening_balance"`
	Tags               string    `json:"tags"`
}
