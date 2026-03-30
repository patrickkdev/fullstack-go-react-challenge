package db

import (
	"api/internal/domain"

	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(dbConn *gorm.DB) *JobRepository {
	return &JobRepository{db: dbConn}
}

func (r *JobRepository) Create(job domain.Job, ownerID int) (domain.Job, error) {
	model := JobModel{
		Title:       job.Title,
		Description: job.Description,
		Company:     job.Company,
		Location:    job.Location,
		OwnerID:     ownerID,
	}
	if err := r.db.Create(&model).Error; err != nil {
		return domain.Job{}, err
	}
	return domain.Job{ID: model.ID, Title: model.Title, Description: model.Description, Company: model.Company, Location: model.Location, Salary: model.Salary}, nil
}

func (r *JobRepository) ListAll() ([]domain.Job, error) {
	var models []JobModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	res := make([]domain.Job, 0, len(models))
	for _, m := range models {
		res = append(res, domain.Job{ID: m.ID, Title: m.Title, Description: m.Description, Company: m.Company, Location: m.Location, Salary: m.Salary})
	}
	return res, nil
}

func (r *JobRepository) ListByOwner(ownerID int) ([]domain.Job, error) {
	var models []JobModel
	if err := r.db.Where("owner_id = ?", ownerID).Find(&models).Error; err != nil {
		return nil, err
	}
	res := make([]domain.Job, 0, len(models))
	for _, m := range models {
		res = append(res, domain.Job{ID: m.ID, Title: m.Title, Description: m.Description, Company: m.Company, Location: m.Location, Salary: m.Salary})
	}
	return res, nil
}

func (r *JobRepository) GetByID(id int) (domain.Job, error) {
	var m JobModel
	if err := r.db.First(&m, id).Error; err != nil {
		return domain.Job{}, err
	}
	return domain.Job{ID: m.ID, Title: m.Title, Description: m.Description, Company: m.Company, Location: m.Location, Salary: m.Salary}, nil
}
