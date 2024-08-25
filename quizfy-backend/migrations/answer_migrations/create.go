package answer_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/result"
)

func Create(a []result.Answer) error {
	return db.Conn.Create(a).Error
}
