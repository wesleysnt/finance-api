package models

import (
	"github.com/wesleysnt/finance-api/pkg/auth"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Email    string  `json:"email" gorm:"unique"`
	Password *string `json:"password"`
	Currency string  `json:"currency"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != nil {
		hashed, err := auth.HashPassword(*u.Password)

		u.Password = &hashed
		return err
	}

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password != nil {
		hashed, err := auth.HashPassword(*u.Password)

		u.Password = &hashed
		return err
	}

	return
}
