package language

import "github.com/google/uuid"

// Соответсвует таблице user_languages
type UserLang struct {
	Id               uuid.UUID `json:"id"`
	UserEmail        string    `json:"user_email"`
	LanguageId       uuid.UUID `json:"language_id"`
	ProficiencyLevel int       `json:"proficiency_level"`
}

// Соответсвует таблицу language
type Lang struct {
	Id       uuid.UUID `json:"id"`
	Language string    `json:"language"`
}

// UserLanguageResponse – ответ, содержащий ID записи user_languages.
type UserLanguageResponse struct {
	ID uuid.UUID `json:"id"`
}
