package domain

import "time"

type JobApplication struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	JobID     int       `json:"jobId"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
