package db

import "time"

type JobApplicationModel struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	UserID    int
	JobID     int
	Status    string    `gorm:"size:50"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
