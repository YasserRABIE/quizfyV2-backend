package quiz_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
)

func DeleteByID(id uint) error {
	err := db.Conn.Delete(&quiz.Quiz{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
