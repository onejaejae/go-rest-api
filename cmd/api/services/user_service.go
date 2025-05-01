package services

import (
	"fmt"
	"go-rest-api/cmd/api/repositories"
	"go-rest-api/cmd/api/requests"
	"go-rest-api/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	repository *repositories.UserRepository
}

func newUserService(db *gorm.DB) *UserService {
	return &UserService{
		repository: repositories.NewUserRepository(db),
	}
}

func (u *UserService) RegisterUser(req requests.RegisterRequest) error {
	panic("not implemented")
}

func (u *UserService) GetUserByEmail(email string) (*models.User, error) {

	result, err := u.repository.FindByEmail(email)
	fmt.Println("result", result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
