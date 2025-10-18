package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleysnt/finance-api/app/http/models"
	"github.com/wesleysnt/finance-api/pkg/auth"
	"gorm.io/gorm"
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

func TestAccountRepository_GetAccountList(t *testing.T) {
	orm := setupTestDB(t)
	userRepo := NewUserRepository(orm)
	accountRepo := NewAccountRepository(orm)
	ctx := context.Background()
	password := "password"
	users := []models.User{
		models.User{
			Model:    gorm.Model{ID: 1},
			Name:     "user 1",
			Email:    "user1@gmail.com",
			Password: &password,
			Currency: "IDR",
		},
		models.User{
			Model:    gorm.Model{ID: 2},
			Name:     "user 2",
			Email:    "user2@gmail.com",
			Password: &password,
			Currency: "IDR",
		},
	}

	accounts := []models.Account{
		models.Account{
			UserId:      1,
			Name:        "Bank A",
			AccountType: "Saving",
			Balance:     1000000,
			Currency:    "IDR",
			IsActive:    true,
		},
		models.Account{
			UserId:      1,
			Name:        "Cash Money",
			AccountType: "Cash",
			Balance:     250000,
			Currency:    "IDR",
			IsActive:    true,
		},
		models.Account{
			UserId:      2,
			Name:        "Bank",
			AccountType: "Cash",
			Balance:     2000000,
			Currency:    "IDR",
			IsActive:    true,
		},
	}

	for _, v := range users {
		userRepo.CreateUser(&v, ctx)
	}

	for _, v := range accounts {
		accountRepo.CreateAccount(&v, ctx)
	}

	res, err := accountRepo.GetAccountList(users[0].ID, ctx)

	assert.NoError(t, err)

	assert.Len(t, res, 2)
	assert.Equal(t, res[0].Name, "Bank A")
	assert.Equal(t, res[1].Name, "Cash Money")
}
