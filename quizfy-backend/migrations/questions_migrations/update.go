package question_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
)

func UpdateImage(questionID uint, imgPath string) error {
	return db.Conn.Model(&quiz.Question{}).Where("id = ?", questionID).Update("image_path", imgPath).Error
}
