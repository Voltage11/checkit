package app_error

import "net/http"

type ErrorType string

const (
	NotFound     ErrorType = "NOT_FOUND"
	BadRequest   ErrorType = "BAD_REQUEST"
	Conflict     ErrorType = "CONFLICT"
	Unauthorized ErrorType = "UNAUTHORIZED"
	Internal     ErrorType = "INTERNAL_ERROR"
	//Validation   ErrorType = "VALIDATION_ERROR" Пока не использую, детали не нужны при не верном запросе
)

type AppError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Code    int       `json:"-"`
	Err     error     `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) IsNotfound() bool {
	return e.Type == NotFound
}

func (e *AppError) IsBadRequest() bool {
	return e.Type == BadRequest
}

// NewNotFound Конструкторы ошибок
func NewNotFound(message string) *AppError {
	return &AppError{
		Type:    NotFound,
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewBadRequest(message string) *AppError {
	return &AppError{
		Type:    BadRequest,
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestWithError(message string, err error) *AppError {
	return &AppError{
		Type:    BadRequest,
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     err,
	}
}

func NewConflict(message string) *AppError {
	return &AppError{
		Type:    Conflict,
		Message: message,
		Code:    http.StatusConflict,
	}
}

func NewInternal(err error) *AppError {
	return &AppError{
		Type:    Internal,
		Message: "Внутренняя ошибка сервера",
		Code:    http.StatusInternalServerError,
		Err:     err,
	}
}

func NewUnauthorized() *AppError {
	return &AppError{
		Type:    Unauthorized,
		Message: "Неавторизованный доступ",
		Code:    http.StatusUnauthorized,
		Err:     nil,
	}
}

func NewForbidden() *AppError {
	return &AppError{
		Type:    Unauthorized,
		Message: "Доступ запрещен",
		Code:    http.StatusForbidden,
		Err:     nil,
	}
}

// NewValidation Новый конструктор для ошибок валидации
//func NewValidation(message string) *AppError {
//	return &AppError{
//		Type:    Validation,
//		Message: message,
//		Code:    http.StatusBadRequest,
//	}
