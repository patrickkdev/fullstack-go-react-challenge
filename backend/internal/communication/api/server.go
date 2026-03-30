package api

import (
	"api/internal/application"

	"github.com/gin-gonic/gin"
)

func NewRouter(authService *application.AuthService, userService *application.UserService, jobService *application.JobService, appService *application.JobApplicationService) *gin.Engine {
	authController := NewAuthController(authService, userService)
	jobController := NewJobController(jobService)
	jobAppController := NewJobApplicationController(appService)

	r := gin.Default()

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	protected := r.Group("")
	protected.Use(AuthMiddleware(authService))
	{
		protected.GET("/me", authController.Me)
		protected.POST("/jobs", jobController.Create)
		protected.GET("/jobs/mine", jobController.MyJobs)
		protected.POST("/jobs/:id/apply", jobAppController.Apply)
		protected.GET("/applications", jobAppController.ListByUser)
	}

	r.GET("/jobs", jobController.List)

	return r
}
