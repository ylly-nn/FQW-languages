package language

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

// Handler обрабатывает HTTP-запросы для языков пользователей.
type Handler struct {
	manager *LanguageManager
}

// NewHandler создаёт новый экземпляр Handler.
func NewHandler(manager *LanguageManager) *Handler {
	return &Handler{manager: manager}
}

// GetOrCreateUserLanguage обрабатывает GET /user/language?language=название
// Защищённый маршрут – email извлекается из JWT токена.
// Если язык отсутствует у пользователя, он создаётся автоматически.
func (h *Handler) GetOrCreateUserLanguage(w http.ResponseWriter, r *http.Request) {
	// Извлекаем email из контекста (установлен middleware)
	claims, ok := r.Context().Value("user").(jwt.MapClaims)
	if !ok {
		http.Error(w, "unauthorized: missing user claims", http.StatusUnauthorized)
		return
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		http.Error(w, "unauthorized: email not found in token", http.StatusUnauthorized)
		return
	}

	// Получаем название языка из query параметра
	languageName := r.URL.Query().Get("language")
	if languageName == "" {
		http.Error(w, "missing language parameter", http.StatusBadRequest)
		return
	}

	// Вызываем бизнес-логику
	id, err := h.manager.GetUserLanguageID(email, languageName)
	if err != nil {
		log.Printf("GetOrCreateUserLanguage error: %v", err)
		switch {
		case errors.Is(err, ErrEmptyEmail),
			errors.Is(err, ErrEmptyLanguageName),
			errors.Is(err, ErrLangNotFound):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Формируем ответ
	resp := UserLanguageResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("GetOrCreateUserLanguage encode error: %v", err)
	}
}
