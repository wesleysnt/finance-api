package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleysnt/finance-api/app/http/models"
	"github.com/wesleysnt/finance-api/pkg/auth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.User{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Account{})
	assert.NoError(t, err)

	return db
}

func TestUserRepository_GetUserByEmail_Success(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	password, err := auth.HashPassword("testPassword")

	assert.NoError(t, err)

	user := &models.User{
		Name:     "Test User",
		Email:    "testuser@example.com",
		Password: &password,
		Currency: "USD",
	}

	err = repo.CreateUser(user, nil)
	assert.NoError(t, err)

	retrievedUser, err := repo.GetUserByEmail("testuser@example.com", nil)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, retrievedUser.Email)
	assert.Equal(t, user.Password, retrievedUser.Password)
	assert.Equal(t, user.Currency, retrievedUser.Currency)
}

func TestUserRepository_GetUserByEmail_Error(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	password, err := auth.HashPassword("testPassword")

	assert.NoError(t, err)

	user := &models.User{
		Name:     "Test User",
		Email:    "testuser@example.com",
		Password: &password,
		Currency: "USD",
	}

	err = repo.CreateUser(user, nil)
	assert.NoError(t, err)

	retrievedUser, err := repo.GetUserByEmail("nottheuser@example.com", nil)
	assert.Error(t, err)
	assert.Nil(t, retrievedUser)
}

func TestUserRepository_CreateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	password, err := auth.HashPassword("newPassword")
	assert.NoError(t, err)

	user := &models.User{
		Name:     "New User",
		Email:    "newuser@example.com",
		Password: &password,
		Currency: "USD",
	}

	err = repo.CreateUser(user, nil)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
}

func TestUserRepository_CreateUser_duplicate(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	password, err := auth.HashPassword("newPassword")
	assert.NoError(t, err)

	user := &models.User{
		Name:     "New User",
		Email:    "newuser@example.com",
		Password: &password,
		Currency: "USD",
	}

	repo.CreateUser(user, nil)

	otheruser := &models.User{
		Name:     "other User",
		Email:    "newuser@example.com",
		Password: &password,
		Currency: "USD",
	}
	err = repo.CreateUser(otheruser, nil)
	assert.Error(t, err)
	assert.Zero(t, otheruser.ID)
}
