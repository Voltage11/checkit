package models

import "time"

// Theme Тема тестирования
type Theme struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	PassingScore int       `json:"passing_score"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ThemeCreateRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	PassingScore int    `json:"passing_score"`
}

type ThemeUpdateRequest struct {
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	PassingScore *int    `json:"passing_score,omitempty"`
	IsActive     *bool   `json:"is_active,omitempty"`
}
