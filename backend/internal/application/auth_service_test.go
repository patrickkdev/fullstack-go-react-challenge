package application

import (
	"errors"
	"testing"

	"api/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var ErrNotFound = errors.New("not found")

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id int) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) GetBySessionToken(token string) (domain.User, error) {
	args := m.Called(token)
	return args.Get(0).(domain.User), args.Error(1)
}

func TestAuthService_Register(t *testing.T) {
	t.Run("successful registration", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewAuthService(mockRepo)

		mockRepo.
			On("GetByEmail", "john@example.com").
			Return(domain.User{}, ErrNotFound).
			Once()

		mockRepo.
			On("Create", mock.MatchedBy(func(u domain.User) bool {
				if u.Email != "john@example.com" || u.Name != "John Doe" {
					return false
				}
				return bcrypt.CompareHashAndPassword(u.PasswordHash, []byte("password123")) == nil
			})).
			Return(domain.NewUser("John Doe", "john@example.com", []byte("irrelevant")), nil).
			Once()

		result, err := service.Register("John Doe", "john@example.com", "password123")

		assert.NoError(t, err)
		assert.Equal(t, "John Doe", result.Name)
		assert.Equal(t, "john@example.com", result.Email)

		mockRepo.AssertExpectations(t)
	})

	t.Run("missing required fields", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewAuthService(mockRepo)

		_, err := service.Register("", "john@example.com", "password123")
		assert.Error(t, err)

		_, err = service.Register("John", "", "password123")
		assert.Error(t, err)

		_, err = service.Register("John", "john@example.com", "")
		assert.Error(t, err)
	})

	t.Run("email already registered", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewAuthService(mockRepo)

		existingUser := domain.NewUser("Existing", "john@example.com", []byte("hash"))

		mockRepo.
			On("GetByEmail", "john@example.com").
			Return(existingUser, nil).
			Once()

		_, err := service.Register("John Doe", "john@example.com", "password123")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "already")

		mockRepo.AssertExpectations(t)
	})

	t.Run("repository create failure", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewAuthService(mockRepo)

		mockRepo.
			On("GetByEmail", "john@example.com").
			Return(domain.User{}, ErrNotFound).
			Once()

		mockRepo.
			On("Create", mock.Anything).
			Return(domain.User{}, errors.New("db error")).
			Once()

		_, err := service.Register("John Doe", "john@example.com", "password123")

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestAuthService_Login(t *testing.T) {
	t.Run("successful login", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewAuthService(mockRepo)

		hash, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		assert.NoError(t, err)

		user := domain.NewUser("John Doe", "john@example.com", hash)

		mockRepo.
			On("GetByEmail", "john@example.com").
			Return(user, nil).
			Once()

		result, err := service.Login("john@example.com", "password123")

		assert.NoError(t, err)
		assert.Equal(t, user.ID, result.ID)
		assert.Equal(t, user.Email, result.Email)

		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid email", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewAuthService(mockRepo)

		mockRepo.
			On("GetByEmail", "wrong@example.com").
			Return(domain.User{}, ErrNotFound).
			Once()

		_, err := service.Login("wrong@example.com", "password123")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid")

		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid password", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewAuthService(mockRepo)

		hash, err := bcrypt.GenerateFromPassword([]byte("correct"), bcrypt.DefaultCost)
		assert.NoError(t, err)

		user := domain.NewUser("John", "john@example.com", hash)

		mockRepo.
			On("GetByEmail", "john@example.com").
			Return(user, nil).
			Once()

		_, err = service.Login("john@example.com", "wrongpassword")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid")

		mockRepo.AssertExpectations(t)
	})

	t.Run("repository failure", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewAuthService(mockRepo)

		mockRepo.
			On("GetByEmail", "john@example.com").
			Return(domain.User{}, errors.New("db error")).
			Once()

		_, err := service.Login("john@example.com", "password123")

		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})
}

func TestAuthService_ValidateToken(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewAuthService(mockRepo)

		user := domain.NewUser("John", "john@example.com", []byte("hash"))

		mockRepo.
			On("GetBySessionToken", "validtoken").
			Return(user, nil).
			Once()

		result, err := service.ValidateToken("validtoken")

		assert.NoError(t, err)
		assert.Equal(t, user.ID, result.ID)

		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid token", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewAuthService(mockRepo)

		mockRepo.
			On("GetBySessionToken", "invalidtoken").
			Return(domain.User{}, ErrNotFound).
			Once()

		_, err := service.ValidateToken("invalidtoken")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid")

		mockRepo.AssertExpectations(t)
	})

	t.Run("repository failure", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewAuthService(mockRepo)

		mockRepo.
			On("GetBySessionToken", "token").
			Return(domain.User{}, errors.New("db error")).
			Once()

		_, err := service.ValidateToken("token")

		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})
}
