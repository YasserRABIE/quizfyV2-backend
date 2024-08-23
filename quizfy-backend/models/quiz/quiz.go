package quiz

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	UserID uint `json:"user_id" gorm:"not null"`

	Title       string `json:"title" gorm:"not null" binding:"required"`
	Description string `json:"description" gorm:"not null" binding:"required"`
	Difficulty  string `json:"difficulty" gorm:"not null" binding:"required"`
	Type        string `json:"type" gorm:"not null" binding:"required"`
	IsTimeBased bool   `json:"is_time_based" gorm:"not null" binding:"required"`
	Duration    *int   `json:"duration" binding:"required"`
	OpensAt     string `json:"opens_at" gorm:"not null" binding:"required"`
	ClosesAt    string `json:"closes_at" gorm:"not null" binding:"required"`

	//? later we can add these fields
	QuestionsCount int `json:"questions_count" gorm:"default:0"`
	TotalDegree    int `json:"total_degree" gorm:"default:0"`
}

// BeforeDelete hook deletes all questions related to the quiz
func (q *Quiz) BeforeDelete(tx *gorm.DB) (err error) {
	var questionIDs []uint

	// Retrieve all question IDs associated with this quiz
	// so we can delete them and their options in a single transaction
	if err := tx.Model(&Question{}).Where("quiz_id = ?", q.ID).Pluck("id", &questionIDs).Error; err != nil {
		return err
	}

	// Delete all questions associated with this quiz and their options will be deleted automatically
	for _, id := range questionIDs {
		if err := tx.Delete(&Question{Model: gorm.Model{ID: id}, QuizID: q.ID}).Error; err != nil {
			return err
		}
	}

	return nil
}

// TableName sets the table name for the Quiz model
func (Quiz) TableName() string {
	return "quizzes"
}
