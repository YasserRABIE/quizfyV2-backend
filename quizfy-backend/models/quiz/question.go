package quiz

import "gorm.io/gorm"

type QuestionReq struct {
	Question
	*ImageData
	NoImage bool `json:"no_image"`
}

type Question struct {
	gorm.Model
	Title      string   `json:"title" gorm:"not null" binding:"required"`
	Difficulty string   `json:"difficulty" gorm:"not null" binding:"required"`
	Type       string   `json:"type" gorm:"not null" binding:"required"` // MCQ | BOOL
	Degree     int      `json:"degree" gorm:"not null" binding:"required"`
	Options    []Option `json:"options" gorm:"foreignKey:QuestionID"`
	BoolAnswer *bool    `json:"bool_answer"` //  Bool
	ImagePath  string   `json:"image_path"`

	QuizID uint `json:"quiz_id" gorm:"not null" binding:"required"`
}

type Option struct {
	gorm.Model
	Title      string `json:"title"`
	IsCorrect  bool   `json:"is_correct"`
	QuestionID uint   `json:"-"`
}

type ImageData struct {
	Image     string `json:"image"`
	Extension string `json:"extension"`
}
