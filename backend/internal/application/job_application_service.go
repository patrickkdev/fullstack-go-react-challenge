package application

import (
	"errors"

	"api/internal/domain"
	"api/internal/infrastructure/db"
)

type JobApplicationService struct {
	jobAppRepo *db.JobApplicationRepository
}

func NewJobApplicationService(jobAppRepo *db.JobApplicationRepository) *JobApplicationService {
	return &JobApplicationService{jobAppRepo: jobAppRepo}
}

func (s *JobApplicationService) Apply(userID, jobID int) (domain.JobApplication, error) {
	if userID == 0 || jobID == 0 {
		return domain.JobApplication{}, errors.New("invalid user or job")
	}
	if exists, _ := s.jobAppRepo.Exists(userID, jobID); exists {
		return domain.JobApplication{}, errors.New("already applied")
	}
	return s.jobAppRepo.Create(domain.JobApplication{UserID: userID, JobID: jobID, Status: "pending"})
}

func (s *JobApplicationService) ListByUser(userID int) ([]domain.JobApplication, error) {
	return s.jobAppRepo.ListByUser(userID)
}
