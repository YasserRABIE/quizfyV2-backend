package result

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	Answer     json.RawMessage `json:"answer" gorm:"not null" binding:"required"`
	SessionID  uint            `json:"session_id" gorm:"not null" binding:"required"`
	QuestionID uint            `json:"question_id" gorm:"not null" binding:"required"`
}
