package db

import (
	"fmt"

	"github.com/quizfy/api/config"
	"github.com/quizfy/api/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic("failed to connect to db due to: " + err.Error())
	}
}

func InitTables() {
	if err := DB.AutoMigrate(&user.Account{}); err != nil {
		panic("failed to migrate db due to: " + err.Error())
	}
}
