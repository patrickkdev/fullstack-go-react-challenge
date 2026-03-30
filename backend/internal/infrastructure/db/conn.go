package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionConfig struct {
	DBName string
	DBHost string
	DBUser string
	DBPass string
	DBPort string
}

func Connect(cfg ConnectionConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&UserModel{}, &JobModel{}, &JobApplicationModel{}); err != nil {
		return nil, err
	}

	return db, nil
}
