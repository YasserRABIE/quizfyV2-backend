package db

import (
	"fmt"

	"github.com/quizfy/api/config"
	"github.com/quizfy/api/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}
	return nil
}

func InitTables() error {
	tables := []interface{}{
		&user.Student{},
		&user.Quizzer{},
	}
	if err := DB.AutoMigrate(tables...); err != nil {
		return fmt.Errorf("failed to migrate db: %w", err)
	}
	return nil
}
