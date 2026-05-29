package language

import (
	"database/sql"
	"errors"
	"fmt"

	"src/internal/db"

	"github.com/google/uuid"
)

var (
	ErrLangNotFound       = errors.New("lang not found")
	ErrLangByUserNotFound = errors.New("lang by user not found")
)

type LanguageStorage interface {
	GetLanguageByName(nameLanguage string) (*Lang, error)
	GetUserLanguagesByEmail(email string) ([]*UserLang, error) // убрал звёздочку перед []
	AddUserLanguages(email string, lang uuid.UUID) (*UserLang, error)
}

type PostgresLanguageStorage struct {
	*db.Storage
}

func NewPostgresLanguageStorage(sqlDB *sql.DB) *PostgresLanguageStorage {
	return &PostgresLanguageStorage{Storage: db.NewStorage(sqlDB)}
}

// GetByName возвращает язык по его названию.
// Если язык не найден, возвращает nil, ErrLangNotFound.
func (s *PostgresLanguageStorage) GetLanguageByName(name string) (*Lang, error) {
	var lang Lang
	err := s.DB.QueryRow(
		`SELECT id, language FROM languages WHERE language = $1`,
		name,
	).Scan(&lang.Id, &lang.Language)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrLangNotFound
		}
		return nil, ErrLangNotFound
	}
	return &lang, nil
}

// GetUserLanguagesByEmail возвращает все записи о языках пользователя по его email.
// Если записи не найдены, возвращает nil, ErrLangByUserNotFound.
func (s *PostgresLanguageStorage) GetUserLanguagesByEmail(email string) ([]*UserLang, error) {
	rows, err := s.DB.Query(
		`SELECT id, "user", language, proficiency_level 
		 FROM user_languages 
		 WHERE "user" = $1`, email)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var langs []*UserLang
	for rows.Next() {
		var ul UserLang
		if err := rows.Scan(&ul.Id, &ul.UserEmail, &ul.LanguageId, &ul.ProficiencyLevel); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		langs = append(langs, &ul)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	if len(langs) == 0 {
		return nil, ErrLangByUserNotFound
	}
	return langs, nil
}

// AddUserLanguages добавляет запись о языке пользователя с уровнем proficiency_level = 0.
// Возвращает созданную запись.
func (s *PostgresLanguageStorage) AddUserLanguages(email string, lang uuid.UUID) (*UserLang, error) {
	var ul UserLang
	err := s.DB.QueryRow(
		`INSERT INTO user_languages ("user", language, proficiency_level)
		 VALUES ($1, $2, 0)
		 RETURNING id, "user", language, proficiency_level`,
		email, lang,
	).Scan(&ul.Id, &ul.UserEmail, &ul.LanguageId, &ul.ProficiencyLevel)
	if err != nil {
		return nil, fmt.Errorf("insert user_languages failed: %w", err)
	}
	return &ul, nil
}
