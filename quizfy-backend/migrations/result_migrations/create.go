package result_migrations

import (
	"github.com/YasserRABIE/QUIZFYv2/db"
	question_migrations "github.com/YasserRABIE/QUIZFYv2/migrations/questions_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"gorm.io/gorm"
)

func Create(sessionID, quizID uint, result *quiz.Result) error {
	result.QuizID = quizID
	result.SessionID = sessionID
	// validate the answers
	for i := range result.ReviewedAnswers {
		answer := &result.ReviewedAnswers[i]

		var err error
		answer.Question, err = question_migrations.GetByID(answer.QuestionID)
		if err != nil {
			return err
		}
		answer.SessionID = sessionID

		// validate the answer
		switch answer.Question.Type {
		case "MCQ":
			answer.IsCorrect = getCorrectOption(answer.Question.Options) == answer.UserAnswer.OptionID
			updateResultCounts(answer.IsCorrect, result, answer.Question)
		case "BOOL":
			answer.IsCorrect = answer.Question.BoolAnswer == answer.UserAnswer.BoolAnswer
			updateResultCounts(answer.IsCorrect, result, answer.Question)
		}
	}
	return db.Conn.Session(&gorm.Session{SkipHooks: true}).Create(result).Error
}

func updateResultCounts(isCorrect bool, result *quiz.Result, question *quiz.Question) {
	if isCorrect {
		result.CorrectCount++
		result.Score += question.Degree
	}
	result.Total += question.Degree
}

func getCorrectOption(options []quiz.Option) uint {
	for _, option := range options {
		if option.IsCorrect {
			return option.ID
		}
	}
	return 0
}
