package answer

import (
	"github.com/YasserRABIE/QUIZFYv2/migrations/answer_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/result"
)

func GetAll(sessionID uint) ([]result.UserAnswer, error) {
	return answer_migrations.GetAll(sessionID)
}
