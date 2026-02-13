package models

import "time"

type Answer struct {
	ID         int       `json:"id"`
	QuestionID int       `json:"question_id"`
	Text       string    `json:"text"`
	IsCorrect  bool      `json:"is_correct"`
	OrderIndex int       `json:"order_index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AnswerCreateRequest struct {
	Text       string `json:"text"`
	IsCorrect  bool   `json:"is_correct"`
	OrderIndex int    `json:"order_index"`
}

type AnswerUpdateRequest struct {
	Text       *string `json:"text,omitempty"`
	IsCorrect  *bool   `json:"is_correct,omitempty"`
	OrderIndex *int    `json:"order_index,omitempty"`
}
