package quizzer_migration

import (
	"errors"

	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"gorm.io/gorm"
)

func Create(q *user.Quizzer) error {
	err := db.Conn.Create(q).Error
	if err != nil {
		// Check if the error is due to a unique constraint violation
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("this phone number is already registered")
		}
		return errors.New("server error. please try again later")
	}
	return nil
}
