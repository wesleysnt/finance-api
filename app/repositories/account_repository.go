package repositories

import (
	"context"

	"github.com/wesleysnt/finance-api/app/http/models"
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(account *models.Account, ctx context.Context) error
	GetAccountList(userId uint, ctx context.Context) ([]models.Account, error)
}

type accountRepository struct {
	orm *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return accountRepository{
		orm: db,
	}
}

func (r accountRepository) CreateAccount(account *models.Account, ctx context.Context) error {
	if err := r.orm.WithContext(ctx).Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (r accountRepository) GetAccountList(userId uint, ctx context.Context) ([]models.Account, error) {
	var account []models.Account
	err := r.orm.Where("user_id", userId).Find(&account)

	if err.Error != nil {
		return []models.Account{}, err.Error
	}
	return account, nil
}
