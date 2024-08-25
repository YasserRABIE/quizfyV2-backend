package result_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/result"
)

func Get(sessionID uint) (*result.Result, error) {
	var result result.Result
	err := db.Conn.Where("session_id = ?", sessionID).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, err
}
