package services

import "gorm.io/gorm"

type Service struct {
	UserService *UserService
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		UserService: newUserService(db),
	}
}
