package db

type JobApplicationModel struct {
	ID     int `gorm:"primaryKey;autoIncrement"`
	UserID int
	JobID  int
	Status string `gorm:"size:50"`
}
