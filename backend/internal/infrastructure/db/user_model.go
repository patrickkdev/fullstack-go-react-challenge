package db

type UserModel struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"size:255"`
	Email        string `gorm:"size:255;uniqueIndex"`
	PasswordHash []byte `gorm:"size:255"`
	SessionToken string `gorm:"size:255;uniqueIndex"`
}
