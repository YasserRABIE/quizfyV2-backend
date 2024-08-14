package user

import "gorm.io/gorm"

type Quizzer struct {
	gorm.Model
	UserInfo
	Subject       string `json:"subject" gorm:"not null" binding:"required"`
}

type QLoginReq struct {
	Phone    uint   `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}
