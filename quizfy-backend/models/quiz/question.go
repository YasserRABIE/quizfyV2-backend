package quiz

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title         string   `json:"title" gorm:"not null" binding:"required"`
	Difficulty    string   `json:"difficulty" gorm:"not null" binding:"required"`
	Type          string   `json:"type" gorm:"not null" binding:"required"` // MCQ | BOOL
	Degree        int      `json:"degree" gorm:"not null" binding:"required"`
	Options       []Option `json:"options" gorm:"foreignKey:QuestionID"`
	CorrectAnswer *bool    `json:"correct_answer"` //  Bool
}

type Option struct {
	gorm.Model
	Title      string `json:"title" binding:"required"`
	IsCorrect  bool   `json:"is_correct" binding:"required"`
	QuestionID uint   `json:"-"`
}
