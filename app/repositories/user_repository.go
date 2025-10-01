package repositories

import "gorm.io/gorm"

type UserRespository struct {
	orm *gorm.DB
}

func NewAuthRepository() *UserRespository {
	return &UserRespository{}
}
