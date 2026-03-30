package application

import (
	"api/internal/domain"
	"api/internal/infrastructure/db"
)

type UserService struct {
	userRepo *db.UserRepository
}

func NewUserService(userRepo *db.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetByID(id int) (domain.User, error) {
	return s.userRepo.GetByID(id)
}
