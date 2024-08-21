package question_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
)

func Delete(id uint) (string, error) {
	var imagePath string

	err := db.Conn.Model(&quiz.Question{}).
		Where("id = ?", id).
		Pluck("image_path", &imagePath).Error

	if err != nil {
		return "", err
	}

	err = db.Conn.Where("question_id = ?", id).Delete(&quiz.Option{}).Error
	if err != nil {
		return "", err
	}

	err = db.Conn.Delete(&quiz.Question{}, id).Error
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
