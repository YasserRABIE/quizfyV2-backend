package question_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
)

func GetAll(quizID uint) ([]quiz.Question, error) {
	var questions []quiz.Question
	if err := db.Conn.Where("quiz_id = ?", quizID).Preload("Options").Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func GetByID(id uint) (*quiz.Question, error) {
	var question quiz.Question
	if err := db.Conn.Preload("Options").First(&question, id).Error; err != nil {
		return nil, err
	}
	return &question, nil
}
