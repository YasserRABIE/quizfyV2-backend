package user

type UserInfo struct {
	Name          string `json:"name" gorm:"not null" binding:"required"`
	Phone         uint   `json:"phone" gorm:"unique;not null" binding:"required"`
	Password      string `json:"password" gorm:"not null" binding:"required"`
	AcademicStage string `json:"academic_stage" gorm:"not null" binding:"required"`
}
