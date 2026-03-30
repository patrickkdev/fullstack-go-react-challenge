package db

import (
	"api/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) *UserRepository {
	return &UserRepository{db: dbConn}
}

func (r *UserRepository) Create(user domain.User) (domain.User, error) {
	model := UserModel{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		SessionToken: user.SessionToken,
	}
	if err := r.db.Create(&model).Error; err != nil {
		return domain.User{}, err
	}
	return domain.User{ID: model.ID, Name: model.Name, Email: model.Email, PasswordHash: model.PasswordHash, SessionToken: model.SessionToken}, nil
}

func (r *UserRepository) GetByEmail(email string) (domain.User, error) {
	var model UserModel
	if err := r.db.Where("email = ?", email).First(&model).Error; err != nil {
		return domain.User{}, err
	}
	return domain.User{ID: model.ID, Name: model.Name, Email: model.Email, PasswordHash: model.PasswordHash, SessionToken: model.SessionToken}, nil
}

func (r *UserRepository) GetByID(id int) (domain.User, error) {
	var model UserModel
	if err := r.db.First(&model, id).Error; err != nil {
		return domain.User{}, err
	}
	return domain.User{ID: model.ID, Name: model.Name, Email: model.Email, PasswordHash: model.PasswordHash, SessionToken: model.SessionToken}, nil
}

func (r *UserRepository) GetBySessionToken(token string) (domain.User, error) {
	var model UserModel
	if err := r.db.Where("session_token = ?", token).First(&model).Error; err != nil {
		return domain.User{}, err
	}
	return domain.User{ID: model.ID, Name: model.Name, Email: model.Email, PasswordHash: model.PasswordHash, SessionToken: model.SessionToken}, nil
}
