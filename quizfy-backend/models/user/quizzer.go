package user

import "gorm.io/gorm"

type Quizzer struct {
	gorm.Model
	UserInfo
	Subject string `json:"subject" gorm:"not null" binding:"required"`
}
