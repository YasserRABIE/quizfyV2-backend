package quiz_migrations

import (
	"errors"

	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"gorm.io/gorm"
)

// GetAll get all quizzes
func GetAll(userID uint) ([]quiz.Quiz, error) {
	var quizzes []quiz.Quiz

	if err := db.Conn.Where("user_id = ?", userID).Find(&quizzes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("لم يتم العثور على أي اختبارات ليك")
		}
		return nil, errors.New("حدث خطأ أثناء استرجاع اختباراتك")
	}

	return quizzes, nil
}

// GetByID get quiz by id
func GetByID(id uint) (*quiz.Quiz, error) {
	var q quiz.Quiz

	if err := db.Conn.Where("id = ?", id).First(&q).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("لم يتم العثور على الاختبار")
		}
		return nil, errors.New("حدث خطأ أثناء استرجاع الاختبار")
	}

	return &q, nil
}

// ========================================
// quizzer
// ========================================

// GetByQuizzerID get quiz by quizzer id and state
func GetExamsByQuizzerID(quizzerID uint, state string) ([]quiz.Quiz, error) {
	var q []quiz.Quiz
	query := "user_id = ? AND type = 'quiz'"

	switch state {
	case "private":
		query += " AND opens_at::timestamp > NOW()"
	case "public":
		query += " AND opens_at::timestamp <= NOW() AND closes_at::timestamp >= NOW()"
	}

	if err := db.Conn.Where(query, quizzerID).Find(&q).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("لم يتم العثور على الاختبارات")
		}
		return nil, errors.New("حدث خطأ أثناء استرجاع الاختبارات")
	}

	return q, nil
}
