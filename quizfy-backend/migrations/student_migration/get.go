package student_migration

import (
	"errors"

	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"gorm.io/gorm"
)

func Get(phone uint) (*user.Student, error) {
	var s user.Student
	err := db.Conn.Where("phone = ?", phone).First(&s).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("الرقم ده مش مسجل عندنا")
	}

	if err != nil {
		return nil, errors.New("حصل خطأ في السيرفر. حاول تاني بعد شوية")
	}

	return &s, nil
}
