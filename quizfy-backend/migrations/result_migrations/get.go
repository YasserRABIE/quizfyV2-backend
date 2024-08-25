package result_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
)

func Get(sessionID uint) (*quiz.Result, error) {
	var result quiz.Result
	err := db.Conn.Where("session_id = ?", sessionID).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, err
}
