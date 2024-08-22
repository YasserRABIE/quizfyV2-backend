package quiz_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"gorm.io/gorm"
)

func DeleteByID(id uint) error {
	err := db.Conn.Delete(&quiz.Quiz{Model: gorm.Model{ID: id}}).Error
	if err != nil {
		return err
	}
	return nil
}
