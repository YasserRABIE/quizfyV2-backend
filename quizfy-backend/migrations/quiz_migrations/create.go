package quiz_migrations

import (
	"errors"

	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
)

func Create(quizData *quiz.Quiz) error {
	if err := db.Conn.Create(quizData).Error; err != nil {
		return errors.New("لم يتم إنشاء الاختبار بنجاح!، يرجى المحاولة مرة أخرى")
	}
	return nil
}
