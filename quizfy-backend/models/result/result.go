package result

import (
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	SessionID       uint             `json:"session_id" gorm:"not null;uniqueIndex"`
	QuizID          uint             `json:"quiz_id" gorm:"not null"`
	Score           uint             `json:"score" gorm:"default:0"`
	Total           uint             `json:"total" gorm:"default:0"`
	CorrectCount    uint             `json:"correct_count" gorm:"default:0"`
	ReviewedAnswers []ReviewedAnswer `json:"reviewed_answers" gorm:"foreignKey:ResultID;references:ID"`
}

type ReviewedAnswer struct {
	gorm.Model
	Question  *quiz.Question `json:"question" gorm:"foreignKey:QuestionID;references:ID"`
	ResultID  uint           `json:"-"`
	IsCorrect bool           `json:"is_correct"`

	// this data is form the params
	SessionID uint `json:"session_id" gorm:"not null"`
	// this data is submitted by the user
	QuestionID uint `json:"question_id" gorm:"not null"`
	UserAnswer
}
