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
		if err := tx.Delete(&Question{Model: gorm.Model{ID: id}}).Error; err != nil {
			return err
		}
	}

	return nil
}

// TableName sets the table name for the Quiz model
func (Quiz) TableName() string {
	return "quizzes"
}
