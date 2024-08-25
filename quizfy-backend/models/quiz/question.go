package quiz

import (
	"fmt"

	"gorm.io/gorm"
)

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
	Degree     uint     `json:"degree" gorm:"not null" binding:"required"`
	Options    []Option `json:"options" gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE;"`
	BoolAnswer bool     `json:"bool_answer"` //  Bool
	ImagePath  string   `json:"image_path"`
	QuizID     uint     `json:"quiz_id" gorm:"not null" binding:"required"`
}

// AfterCreate hook updates the quiz questions count and total degree
func (q *Question) AfterCreate(tx *gorm.DB) (err error) {
	var quiz Quiz
	if err = tx.First(&quiz, q.QuizID).Error; err != nil {
		return err
	}
	quiz.QuestionsCount++
	quiz.TotalDegree += int(q.Degree)
	if err = tx.Save(&quiz).Error; err != nil {
		return err
	}
	return
}

// BeforeDelete hook deletes all options related to the question
func (q *Question) BeforeDelete(tx *gorm.DB) (err error) {
	// Update the quiz questions count and total degree
	var quiz Quiz
	if err = tx.First(&quiz, q.QuizID).Error; err != nil {
		return err
	}
	fmt.Println("Degree", q.Degree)
	quiz.QuestionsCount--
	quiz.TotalDegree -= int(q.Degree)
	if err = tx.Save(&quiz).Error; err != nil {
		return err
	}

	// Delete all options associated with this question
	tx.Where("question_id = ?", q.ID).Delete(&Option{})
	return
}

// BeforeUpdate hook updates the quiz total degree
func (q *Question) BeforeUpdate(tx *gorm.DB) (err error) {

	// Get the old degree
	var oldDegree uint
	if err = tx.Model(&Question{}).Where("id = ?", q.ID).Select("degree").
		Scan(&oldDegree).Error; err != nil {
		return err
	}

	// Update the quiz total degree
	var quiz Quiz
	if err = tx.First(&quiz, q.QuizID).Error; err != nil {
		return err
	}
	quiz.TotalDegree += int(q.Degree) - int(oldDegree)
	if err = tx.Save(&quiz).Error; err != nil {
		return err
	}
	return
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
