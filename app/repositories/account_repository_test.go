package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleysnt/finance-api/app/http/models"
	"github.com/wesleysnt/finance-api/pkg/auth"
)

func TestAccountRepository_CreateAccount(t *testing.T) {
	orm := setupTestDB(t)
	userRepo := NewUserRepository(orm)
	accountRepo := NewAccountRepository(orm)

	ctx := context.Background()
	hash, _ := auth.HashPassword("testpass")
	user := models.User{
		Name:     "testUser",
		Email:    "testuser@gmail.com",
		Password: &hash,
		Currency: "IDR",
	}

	err := userRepo.CreateUser(&user, ctx)
	assert.NoError(t, err)

	account := models.Account{
		UserId:      user.ID,
		Name:        "test Account",
		AccountType: "test",
		Currency:    "IDR",
	}
	err = accountRepo.CreateAccount(&account, ctx)

	assert.NoError(t, err)
	assert.NotZero(t, account.ID)
}
