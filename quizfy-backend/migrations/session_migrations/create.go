package session_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/session"
)

func Create(quizID, userID uint) (*session.Session, error) {
	session := &session.Session{
		QuizID: quizID,
		UserID: userID,
	}
	if err := db.Conn.Create(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}
