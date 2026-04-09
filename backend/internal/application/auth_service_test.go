package application

import (
	"errors"
	"testing"

	"api/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// MockUserRepository is a mock implementation of UserRepository
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
	mockRepo := new(MockUserRepository)
	service := NewAuthService(mockRepo)

	t.Run("successful registration", func(t *testing.T) {
		user := domain.NewUser("John Doe", "john@example.com", []byte("hashedpass"))
		mockRepo.On("GetByEmail", "john@example.com").Return(domain.User{}, errors.New("not found")).Once()
		mockRepo.On("Create", mock.AnythingOfType("domain.User")).Return(user, nil).Once()

		result, err := service.Register("John Doe", "john@example.com", "password123")

		assert.NoError(t, err)
		assert.Equal(t, "John Doe", result.Name)
		assert.Equal(t, "john@example.com", result.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("missing required fields", func(t *testing.T) {
		_, err := service.Register("", "john@example.com", "password123")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required")

		_, err = service.Register("John", "", "password123")
		assert.Error(t, err)

		_, err = service.Register("John", "john@example.com", "")
		assert.Error(t, err)
	})

	t.Run("email already registered", func(t *testing.T) {
		existingUser := domain.NewUser("Existing", "john@example.com", []byte("hash"))
		mockRepo.On("GetByEmail", "john@example.com").Return(existingUser, nil).Once()

		_, err := service.Register("John Doe", "john@example.com", "password123")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "already registered")
		mockRepo.AssertExpectations(t)
	})
}

func TestAuthService_Login(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewAuthService(mockRepo)

	t.Run("successful login", func(t *testing.T) {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		user := domain.NewUser("John Doe", "john@example.com", hash)
		mockRepo.On("GetByEmail", "john@example.com").Return(user, nil).Once()

		result, err := service.Login("john@example.com", "password123")

		assert.NoError(t, err)
		assert.Equal(t, user.ID, result.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid email", func(t *testing.T) {
		mockRepo.On("GetByEmail", "wrong@example.com").Return(domain.User{}, errors.New("not found")).Once()

		_, err := service.Login("wrong@example.com", "password123")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid credentials")
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid password", func(t *testing.T) {
		hash, _ := bcrypt.GenerateFromPassword([]byte("correct"), bcrypt.DefaultCost)
		user := domain.NewUser("John", "john@example.com", hash)
		mockRepo.On("GetByEmail", "john@example.com").Return(user, nil).Once()

		_, err := service.Login("john@example.com", "wrongpassword")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid credentials")
		mockRepo.AssertExpectations(t)
	})
}

func TestAuthService_ValidateToken(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewAuthService(mockRepo)

	t.Run("valid token", func(t *testing.T) {
		user := domain.NewUser("John", "john@example.com", []byte("hash"))
		mockRepo.On("GetBySessionToken", "validtoken").Return(user, nil).Once()

		result, err := service.ValidateToken("validtoken")

		assert.NoError(t, err)
		assert.Equal(t, user.ID, result.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid token", func(t *testing.T) {
		mockRepo.On("GetBySessionToken", "invalidtoken").Return(domain.User{}, errors.New("not found")).Once()

		_, err := service.ValidateToken("invalidtoken")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid token")
		mockRepo.AssertExpectations(t)
	})
}