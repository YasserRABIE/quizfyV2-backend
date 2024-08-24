package session_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/session"
)

func Get(userID, quizID uint) (*session.Session, error) {
	var s session.Session

	err := db.Conn.Where(&session.Session{QuizID: quizID, UserID: userID}).First(&s).Error
	if err != nil {
		return nil, err
	}

	return &s, nil
}
