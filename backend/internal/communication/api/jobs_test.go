package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockJobService struct {
	mock.Mock
}

func (m *MockJobService) ListAll() ([]domain.Job, error) {
	args := m.Called()
	return args.Get(0).([]domain.Job), args.Error(1)
}

func (m *MockJobService) Create(job domain.Job, userID int) (domain.Job, error) {
	args := m.Called(job, userID)
	return args.Get(0).(domain.Job), args.Error(1)
}

func (m *MockJobService) ListByOwner(userID int) ([]domain.Job, error) {
	args := m.Called(userID)
	return args.Get(0).([]domain.Job), args.Error(1)
}

// --- helpers ---

func setupRouter(controller *JobController) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("/jobs", controller.List)
	r.POST("/jobs", controller.Create)
	r.GET("/jobs/me", controller.MyJobs)

	return r
}

func addUserToContext(user domain.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user", user)
		c.Next()
	}
}

// --- tests ---

func TestJobController_List(t *testing.T) {
	t.Run("returns all jobs", func(t *testing.T) {
		mockService := new(MockJobService)
		controller := NewJobController(mockService)
		router := setupRouter(controller)

		jobs := []domain.Job{
			{Title: "Backend Dev", Description: "Go"},
		}

		mockService.On("ListAll").Return(jobs, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/jobs", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Backend Dev")

		mockService.AssertExpectations(t)
	})

	t.Run("filters jobs by search", func(t *testing.T) {
		mockService := new(MockJobService)
		controller := NewJobController(mockService)
		router := setupRouter(controller)

		jobs := []domain.Job{
			{Title: "Backend Dev", Description: "Go"},
			{Title: "Frontend Dev", Description: "React"},
		}

		mockService.On("ListAll").Return(jobs, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/jobs?search=go", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Backend Dev")
		assert.NotContains(t, w.Body.String(), "Frontend Dev")

		mockService.AssertExpectations(t)
	})

	t.Run("service failure", func(t *testing.T) {
		mockService := new(MockJobService)
		controller := NewJobController(mockService)
		router := setupRouter(controller)

		mockService.On("ListAll").Return([]domain.Job{}, assert.AnError).Once()

		req := httptest.NewRequest(http.MethodGet, "/jobs", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		mockService.AssertExpectations(t)
	})
}

func TestJobController_Create(t *testing.T) {
	t.Run("successful creation", func(t *testing.T) {
		mockService := new(MockJobService)

		gin.SetMode(gin.TestMode)
		r := gin.New()

		controller := NewJobController(mockService)
		user := domain.User{ID: 1}

		r.POST("/jobs",
			addUserToContext(user),
			controller.Create,
		)

		payload := `{
			"title": "Backend Dev",
			"description": "Go dev"
		}`

		mockService.
			On("Create", mock.MatchedBy(func(j domain.Job) bool {
				return j.Title == "Backend Dev" &&
					j.Description == "Go dev"
			}), 1).
			Return(domain.Job{Title: "Backend Dev"}, nil).
			Once()

		req := httptest.NewRequest(http.MethodPost, "/jobs", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "Backend Dev")

		mockService.AssertExpectations(t)
	})

	t.Run("invalid payload", func(t *testing.T) {
		mockService := new(MockJobService)
		controller := NewJobController(mockService)
		router := setupRouter(controller)

		req := httptest.NewRequest(http.MethodPost, "/jobs", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("unauthorized", func(t *testing.T) {
		mockService := new(MockJobService)
		controller := NewJobController(mockService)
		router := setupRouter(controller)

		payload := `{"title":"x","description":"y"}`

		req := httptest.NewRequest(http.MethodPost, "/jobs", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("service failure", func(t *testing.T) {
		mockService := new(MockJobService)

		gin.SetMode(gin.TestMode)
		r := gin.New()

		controller := NewJobController(mockService)
		user := domain.User{ID: 1}

		r.POST("/jobs",
			addUserToContext(user),
			controller.Create,
		)

		payload := `{
			"title": "Backend Dev",
			"description": "Go dev"
		}`

		mockService.
			On("Create", mock.Anything, 1).
			Return(domain.Job{}, assert.AnError).
			Once()

		req := httptest.NewRequest(http.MethodPost, "/jobs", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		mockService.AssertExpectations(t)
	})
}

func TestJobController_MyJobs(t *testing.T) {
	t.Run("returns user jobs", func(t *testing.T) {
		mockService := new(MockJobService)

		gin.SetMode(gin.TestMode)
		r := gin.New()

		controller := NewJobController(mockService)
		user := domain.User{ID: 1}

		r.GET("/jobs/me",
			addUserToContext(user),
			controller.MyJobs,
		)

		jobs := []domain.Job{
			{Title: "My Job"},
		}

		mockService.
			On("ListByOwner", 1).
			Return(jobs, nil).
			Once()

		req := httptest.NewRequest(http.MethodGet, "/jobs/me", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "My Job")

		mockService.AssertExpectations(t)
	})

	t.Run("unauthorized", func(t *testing.T) {
		mockService := new(MockJobService)
		controller := NewJobController(mockService)
		router := setupRouter(controller)

		req := httptest.NewRequest(http.MethodGet, "/jobs/me", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("service failure", func(t *testing.T) {
		mockService := new(MockJobService)

		gin.SetMode(gin.TestMode)
		r := gin.New()

		controller := NewJobController(mockService)
		user := domain.User{ID: 1}

		r.GET("/jobs/me",
			addUserToContext(user),
			controller.MyJobs,
		)

		mockService.
			On("ListByOwner", 1).
			Return([]domain.Job{}, assert.AnError).
			Once()

		req := httptest.NewRequest(http.MethodGet, "/jobs/me", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		mockService.AssertExpectations(t)
	})
}
