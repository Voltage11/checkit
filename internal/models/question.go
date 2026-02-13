package models

import "time"

type QuestionType string

const (
	SingleChoice   QuestionType = "single"
	MultipleChoice QuestionType = "multiple"
)

type Question struct {
	ID         int          `json:"id"`
	ThemeID    int          `json:"theme_id"`
	Text       string       `json:"text"`
	Type       QuestionType `json:"type"`
	Points     int          `json:"points"`
	OrderIndex int          `json:"order_index"`
	IsActive   bool         `json:"is_active"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`

	Answers []Answer `json:"answers,omitempty" db:"-"`
}

type QuestionCreateRequest struct {
	ThemeID    int                   `json:"theme_id"`
	Text       string                `json:"text"`
	Type       QuestionType          `json:"type"`
	Points     int                   `json:"points"`
	OrderIndex int                   `json:"order_index"`
	Answers    []AnswerCreateRequest `json:"answers"`
}

type QuestionUpdateRequest struct {
	Text       *string       `json:"text,omitempty"`
	Type       *QuestionType `json:"type,omitempty"`
	Points     *int          `json:"points,omitempty"`
	OrderIndex *int          `json:"order_index,omitempty"`
	IsActive   *bool         `json:"is_active,omitempty"`
}
