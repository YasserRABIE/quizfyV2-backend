package answer_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/result"
)

func GetAll(sessionID uint) ([]result.UserAnswer, error) {
	var answers []result.UserAnswer
	if err := db.Conn.Where("session_id = ?", sessionID).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}
