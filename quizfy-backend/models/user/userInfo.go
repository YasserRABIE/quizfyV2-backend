package user

type UserInfo struct {
	Name          string `json:"name" gorm:"not null"`
	Phone         int    `json:"phone" gorm:"unique;not null"`
	Password      string `json:"password" gorm:"not null"`
	AcademicStage string `json:"academic_stage" gorm:"not null"`
}
