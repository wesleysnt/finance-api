package repositories

import (
	"context"

	"github.com/wesleysnt/finance-api/app/http/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Define your methods here, e.g.:
	GetUserByEmail(email string, ctx context.Context) (*models.User, error)
	CreateUser(user *models.User, ctx context.Context) error
}

type userRepository struct {
	orm *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		orm: db,
	}
}
func (r *userRepository) GetUserByEmail(email string, ctx context.Context) (*models.User, error) {
	var user models.User
	if err := r.orm.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *models.User, ctx context.Context) error {
	if err := r.orm.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}
