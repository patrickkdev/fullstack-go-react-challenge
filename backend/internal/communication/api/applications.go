package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"api/internal/application"
)

type JobApplicationController struct {
	service *application.JobApplicationService
}

func NewJobApplicationController(service *application.JobApplicationService) *JobApplicationController {
	return &JobApplicationController{service: service}
}

func (j *JobApplicationController) Apply(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil || jobID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	user, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	application, err := j.service.Apply(user.ID, jobID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, application)
}

func (j *JobApplicationController) ListByUser(c *gin.Context) {
	user, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	applications, err := j.service.ListByUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, applications)
}
