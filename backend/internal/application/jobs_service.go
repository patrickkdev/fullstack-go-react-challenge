package application

import (
	"api/internal/domain"
	"api/internal/infrastructure/db"
)

type JobService struct {
	jobRepo *db.JobRepository
}

func NewJobService(jobRepo *db.JobRepository) *JobService {
	return &JobService{jobRepo: jobRepo}
}

func (s *JobService) Create(job domain.Job, ownerID int) (domain.Job, error) {
	return s.jobRepo.Create(job, ownerID)
}

func (s *JobService) ListAll() ([]domain.Job, error) {
	return s.jobRepo.ListAll()
}

func (s *JobService) ListByOwner(ownerID int) ([]domain.Job, error) {
	return s.jobRepo.ListByOwner(ownerID)
}
