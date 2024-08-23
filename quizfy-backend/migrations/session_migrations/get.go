package session_migrations

import (
	"errors"

	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/session"
	"gorm.io/gorm"
)

func Get(userID, quizID uint) (*session.Session, error) {
	var s session.Session

	err := db.Conn.Where(&session.Session{QuizID: quizID, UserID: userID}).First(&s).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Create(userID, quizID)
		}
		return nil, err
	}

	return &s, nil
}
