package user

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name          string `json:"name" gorm:"not null" binding:"required"`
	Phone         string `json:"phone" gorm:"unique;not null" binding:"required"`
	Password      string `json:"password" gorm:"not null" binding:"required"`
	AcademicLevel string `json:"academic_level" gorm:"not null" binding:"required"`
	AcademicYear  string `json:"academic_year" gorm:"not null" binding:"required"`
	AccountType   string `json:"account_type" gorm:"not null" binding:"required"`
	Subject       string `json:"subject" gorm:"not null" binding:"required"`
}

type LoginReq struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}
