package user

type Account struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Subject  string `json:"subject"`
	Password string `json:"password"`
}
