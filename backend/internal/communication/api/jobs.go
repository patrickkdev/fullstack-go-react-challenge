package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"api/internal/domain"
)

type JobService interface {
	ListAll() ([]domain.Job, error)
	Create(job domain.Job, userID int) (domain.Job, error)
	ListByOwner(userID int) ([]domain.Job, error)
}

type JobController struct {
	service JobService
}

func NewJobController(service JobService) *JobController {
	return &JobController{service: service}
}

func (j *JobController) List(c *gin.Context) {
	search := c.Query("search")
	jobs, err := j.service.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if search == "" {
		c.JSON(http.StatusOK, jobs)
		return
	}

	filtered := make([]domain.Job, 0, len(jobs))
	for _, job := range jobs {
		if containsIgnoreCase(job.Title, search) || containsIgnoreCase(job.Description, search) {
			filtered = append(filtered, job)
		}
	}

	c.JSON(http.StatusOK, filtered)
}

func (j *JobController) Create(c *gin.Context) {
	var payload struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Company     string `json:"company"`
		Location    string `json:"location"`
		Salary      string `json:"salary"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	job, err := j.service.Create(domain.Job{
		Title:       payload.Title,
		Description: payload.Description,
		Company:     payload.Company,
		Location:    payload.Location,
		Salary:      payload.Salary,
	}, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, job)
}

func (j *JobController) MyJobs(c *gin.Context) {
	user, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	jobs, err := j.service.ListByOwner(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func containsIgnoreCase(value, substr string) bool {
	return strings.Contains(strings.ToLower(value), strings.ToLower(substr))
}
