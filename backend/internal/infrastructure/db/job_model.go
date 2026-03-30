package db

type JobModel struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"size:255"`
	Description string `gorm:"type:text"`
	Company     string `gorm:"size:255"`
	Location    string `gorm:"size:255"`
	OwnerID     int
	Salary      string
}
