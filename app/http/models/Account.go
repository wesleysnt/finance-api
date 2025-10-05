package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserId      uint    `json:"user_id"`
	Name        string  `json:"name"`
	AccountType string  `json:"account_type"`
	Balance     float64 `json:"balance"`
	Currency    string  `json:"currency"`
	IsActive    bool    `json:"is_active"`
}
