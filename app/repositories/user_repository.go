package repositories

import (
	"context"

	"github.com/wesleysnt/finance-api/app/http/models"
	"github.com/wesleysnt/finance-api/pkg"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Define your methods here, e.g.:
	GetUserByEmail(email string, ctx context.Context) (*models.User, error)
}

type userRepository struct {
	orm *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		orm: pkg.Orm(),
	}
}
func (r *userRepository) GetUserByEmail(email string, ctx context.Context) (*models.User, error) {
	var user models.User
	if err := r.orm.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
