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
