package quiz_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
)

func Update(id uint, quizData *quiz.Quiz) error {
	err := db.Conn.Model(&quiz.Quiz{}).Where("id = ?", id).Updates(
		quiz.Quiz{
			Title:       quizData.Title,
			Description: quizData.Description,
			Difficulty:  quizData.Difficulty,
			Type:        quizData.Type,
			IsTimeBased: quizData.IsTimeBased,
			Duration:    quizData.Duration,
			OpensAt:     quizData.OpensAt,
			ClosesAt:    quizData.ClosesAt,
		},
	).Error
	if err != nil {
		return err
	}
	return nil
}
