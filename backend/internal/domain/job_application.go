package domain

type JobApplication struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	JobID  int    `json:"jobId"`
	Status string `json:"status"`
}
