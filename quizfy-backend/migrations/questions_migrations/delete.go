package question_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"gorm.io/gorm"
)

func DeleteByID(id uint) (string, error) {
	var q struct {
		ImagePath string
		QuizID    uint
		Degree    uint
	}

	// Get the image path and quiz ID before deletion
	err := db.Conn.Model(&quiz.Question{}).
		Where("id = ?", id).
		Select("image_path, quiz_id, degree").
		Scan(&q).Error
	if err != nil {
		return "", err
	}

	// Delete the question
	err = db.Conn.Delete(&quiz.Question{Model: gorm.Model{ID: id}, QuizID: q.QuizID, Degree: q.Degree}).Error
	if err != nil {
		return "", err
	}

	return q.ImagePath, nil
}

func DeleteOptions(questionID uint, excludeIDs []uint) error {
	query := db.Conn.Where("question_id = ?", questionID)

	if len(excludeIDs) > 0 {
		query = query.Where("id NOT IN (?)", excludeIDs)
	}

	if err := query.Delete(&quiz.Option{}).Error; err != nil {
		return err
	}

	return nil
}
