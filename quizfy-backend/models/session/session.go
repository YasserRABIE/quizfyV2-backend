package session

import (
	"time"

	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"gorm.io/gorm"
)

type SessionStatus string

const (
	Active   SessionStatus = "active"
	Reviewed SessionStatus = "reviewed"
)

type Session struct {
	gorm.Model
	QuizID uint `json:"quiz_id" gorm:"not null;uniqueIndex:idx_user_quiz" binding:"required"`
	UserID uint `json:"user_id" gorm:"not null;uniqueIndex:idx_user_quiz" binding:"required"`

	Status    SessionStatus `json:"status" gorm:"default:'active'"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
}

// BeforeCreate hook sets the start and end time of the session
func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	// Get the quiz duration
	var q quiz.Quiz
	if err = tx.First(&q, s.QuizID).Error; err != nil {
		return err
	}
	if !q.IsTimeBased {
		return
	}
	s.StartTime = time.Now()
	s.EndTime = s.StartTime.Add(time.Duration(*q.Duration) * time.Minute)
	return
}

// BeforeUpdate hook updates the start time of the session
// and checks if the session is expired
func (s *Session) BeforeUpdate(tx *gorm.DB) (err error) {
	// Check if the session is expired
	if s.EndTime.Before(time.Now()) {
		s.Status = Reviewed
	}
	s.StartTime = time.Now()
	return
}
