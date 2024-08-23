package db

import (
	"fmt"

	"github.com/YasserRABIE/QUIZFYv2/config"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"github.com/YasserRABIE/QUIZFYv2/models/session"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort)

	var err error
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		TranslateError:         true,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}
	return nil
}

func InitTables() error {
	tables := []interface{}{
		&user.Account{},
		&quiz.Quiz{},
		&quiz.Question{},
		&quiz.Option{},
		&session.Session{},
	}
	if err := Conn.AutoMigrate(tables...); err != nil {
		return fmt.Errorf("failed to migrate db: %w", err)
	}
	return nil
}
