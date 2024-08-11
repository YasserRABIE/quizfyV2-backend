package user

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	UserInfo
	AcademicYear string `json:"academic_year" gorm:"not null" binding:"required"`
}
