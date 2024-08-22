package question_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"gorm.io/gorm"
)

func DeleteByID(id uint) (string, error) {
	var imagePath string

	// Get the image path before deletion
	err := db.Conn.Model(&quiz.Question{}).
		Where("id = ?", id).
		Pluck("image_path", &imagePath).Error
	if err != nil {
		return "", err
	}

	// Delete the question, related options will be deleted automatically
	err = db.Conn.Delete(&quiz.Question{Model: gorm.Model{ID: id}}).Error
	if err != nil {
		return "", err
	}

	return imagePath, nil
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
