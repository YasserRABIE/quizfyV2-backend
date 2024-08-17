package user_migrations

import (
	"errors"

	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"gorm.io/gorm"
)

func GetByPhone(phone string) (*user.Account, error) {
	var a user.Account
	err := db.Conn.Where("phone = ?", phone).First(&a).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("الرقم ده مش مسجل عندنا")
	}

	if err != nil {
		return nil, errors.New("حصل خطأ في السيرفر. حاول تاني بعد شوية")
	}

	return &a, nil
}

func GetById(id uint) (*user.Account, error) {
	var a user.Account
	err := db.Conn.Where("id = ?", id).First(&a).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("الحساب ده مش موجود")
	}

	if err != nil {
		return nil, errors.New("حصل خطأ في السيرفر. حاول تاني بعد شوية")
	}

	return &a, nil
}
