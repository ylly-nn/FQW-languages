package language

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrEmptyEmail           = errors.New("email cannot be empty")
	ErrEmptyLanguageName    = errors.New("language name cannot be empty")
	ErrUserLanguageNotFound = errors.New("user language not found")
)

// LanguageManager содержит бизнес-логику работы с языками пользователей.
type LanguageManager struct {
	storage LanguageStorage
}

// NewLanguageManager создаёт новый экземпляр LanguageManager.
func NewLanguageManager(s LanguageStorage) *LanguageManager {
	return &LanguageManager{storage: s}
}

func (m *LanguageManager) GetUserLanguageID(email string, languageName string) (uuid.UUID, error) {
	if email == "" {
		return uuid.Nil, ErrEmptyEmail
	}
	if languageName == "" {
		return uuid.Nil, ErrEmptyLanguageName
	}

	// Приведение к нижнему регистру
	languageName = strings.ToLower(languageName)

	// 1. Ищем язык по названию
	lang, err := m.storage.GetLanguageByName(languageName)
	if err != nil {
		return uuid.Nil, err
	}

	// 2. Получаем все языки пользователя
	userLangs, err := m.storage.GetUserLanguagesByEmail(email)
	if err != nil {
		// Если пользователь вообще не найден — создаём запись
		if errors.Is(err, ErrLangByUserNotFound) {
			newUL, createErr := m.storage.AddUserLanguages(email, lang.Id)
			if createErr != nil {
				return uuid.Nil, fmt.Errorf("failed to create user language: %w", createErr)
			}
			return newUL.Id, nil
		}
		return uuid.Nil, err
	}

	// 3. Ищем нужный язык в списке
	for _, ul := range userLangs {
		if ul.LanguageId == lang.Id {
			return ul.Id, nil
		}
	}

	// 4. Язык не найден в списке — создаём
	newUL, err := m.storage.AddUserLanguages(email, lang.Id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create user language: %w", err)
	}
	return newUL.Id, nil
}
