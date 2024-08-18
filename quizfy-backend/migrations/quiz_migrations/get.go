package quiz_migrations

import (
	"errors"

	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"gorm.io/gorm"
)

func GetAll(userID uint) ([]quiz.Quiz, error) {
	var quizzes []quiz.Quiz

	if err := db.Conn.Where("user_id = ?", userID).Find(&quizzes).Error; err != nil {
		// Return specific error if no records found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("لم يتم العثور على أي اختبارات ليك")
		}
		// Return the actual error with context if something else went wrong
		return nil, errors.New("حدث خطأ أثناء استرجاع اختباراتك")
	}

	return quizzes, nil
}
