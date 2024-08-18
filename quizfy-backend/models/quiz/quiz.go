package quiz

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"not null"`

	Title       string `json:"title" gorm:"not null" binding:"required"`
	Description string `json:"description" gorm:"not null" binding:"required"`
	Difficulty  string `json:"difficulty" gorm:"not null" binding:"required"`
	Type        string `json:"type" gorm:"not null" binding:"required"`
	IsTimeBased bool   `json:"is_time_based" gorm:"not null" binding:"required"`
	Duration    *int   `json:"duration" binding:"required"`
	OpensAt     string `json:"opens_at" gorm:"not null" binding:"required"`
	ClosesAt    string `json:"closes_at" gorm:"not null" binding:"required"`
}

// TableName sets the table name for the Quiz model
func (Quiz) TableName() string {
	return "quizzes"
}
