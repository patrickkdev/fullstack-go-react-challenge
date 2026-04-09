package application

import (
	"errors"

	"api/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	GetByEmail(email string) (domain.User, error)
	GetByID(id int) (domain.User, error)
	GetBySessionToken(token string) (domain.User, error)
}

type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(name, email, password string) (domain.User, error) {
	if name == "" || email == "" || password == "" {
		return domain.User{}, errors.New("name, email and password are required")
	}

	_, err := s.userRepo.GetByEmail(email)
	if err == nil {
		return domain.User{}, errors.New("email already registered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}

	user, err := s.userRepo.Create(domain.NewUser(name, email, hash))
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *AuthService) Login(email, password string) (domain.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return domain.User{}, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		return domain.User{}, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *AuthService) ValidateToken(tokenStr string) (domain.User, error) {
	user, err := s.userRepo.GetBySessionToken(tokenStr)
	if err != nil {
		return domain.User{}, errors.New("invalid token")
	}
	return user, nil
}
