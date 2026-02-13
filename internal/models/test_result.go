package models

import "time"

type TestResult struct {
	ID           int        `json:"id"`
	UserID       int        `json:"user_id"`
	ThemeID      int        `json:"theme_id"`
	Score        int        `json:"score"`
	MaxScore     int        `json:"max_score"`
	PassingScore int        `json:"passing_score"`
	IsPassed     bool       `json:"is_passed"`
	StartedAt    time.Time  `json:"started_at"`
	CompletedAt  *time.Time `json:"completed_at"`

	User  *User  `json:"user,omitempty"`
	Theme *Theme `json:"theme,omitempty"`
}

type UserAnswer struct {
	ID           int       `json:"id"`
	TestResultID int       `json:"test_result_id"`
	QuestionID   int       `json:"question_id"`
	AnswerID     int       `json:"answer_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type TestSubmitRequest struct {
	ThemeID int           `json:"theme_id"`
	Answers map[int][]int `json:"answers"` // question_id -> []answer_id
}
