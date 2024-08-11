package quizzer_migration

import (
	"errors"

	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"gorm.io/gorm"
)

func Create(q *user.Quizzer) error {
	err := db.Conn.Create(q).Error
	if err != nil {
		// Check if the error is due to a unique constraint violation
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("الرقم ده مسجل قبل كده")
		}
		return errors.New("حصل خطأ في السيرفر. حاول تاني بعد شوية")
	}
	return nil
}
