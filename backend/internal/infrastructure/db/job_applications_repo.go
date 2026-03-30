package db

import (
	"api/internal/domain"

	"gorm.io/gorm"
)

type JobApplicationRepository struct {
	db *gorm.DB
}

func NewJobApplicationRepository(dbConn *gorm.DB) *JobApplicationRepository {
	return &JobApplicationRepository{db: dbConn}
}

func (r *JobApplicationRepository) Create(app domain.JobApplication) (domain.JobApplication, error) {
	model := JobApplicationModel{
		UserID: app.UserID,
		JobID:  app.JobID,
		Status: app.Status,
	}
	if err := r.db.Create(&model).Error; err != nil {
		return domain.JobApplication{}, err
	}
	return domain.JobApplication{ID: model.ID, UserID: model.UserID, JobID: model.JobID, Status: model.Status}, nil
}

func (r *JobApplicationRepository) ListByUser(userID int) ([]domain.JobApplication, error) {
	var models []JobApplicationModel
	if err := r.db.Where("user_id = ?", userID).Find(&models).Error; err != nil {
		return nil, err
	}
	res := make([]domain.JobApplication, 0, len(models))
	for _, m := range models {
		res = append(res, domain.JobApplication{ID: m.ID, UserID: m.UserID, JobID: m.JobID, Status: m.Status})
	}
	return res, nil
}

func (r *JobApplicationRepository) Exists(userID, jobID int) (bool, error) {
	var count int64
	if err := r.db.Model(&JobApplicationModel{}).Where("user_id = ? AND job_id = ?", userID, jobID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
