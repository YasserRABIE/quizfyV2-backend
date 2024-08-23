package session_migrations

import (
	"time"

	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/session"
)

func Update(sessionID uint) (*session.Session, error) {
	// Load the session from the database
	var s session.Session
	if err := db.Conn.First(&s, sessionID).Error; err != nil {
		return nil, err
	}

	if s.Status == session.Active {
		// Update the StartTime
		s.StartTime = time.Now()

		// Check if the session is over
		if s.EndTime.Before(time.Now()) {
			s.Status = session.Reviewed
		}

	}
	// Save the session
	if err := db.Conn.Save(&s).Error; err != nil {
		return nil, err
	}

	return &s, nil
}
