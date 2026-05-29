package swagger

import "src/internal/language"

// GetOrCreateUserLanguage обрабатывает GET /language?language=...
// @Summary      Получить или создать язык пользователя
// @Description  Возвращает ID записи в таблице user_languages для указанного языка. Если язык не найден у пользователя, он автоматически добавляется.
// @Tags         language
// @Accept       json
// @Produce      json
// @Param        language  query     string  true   "Название языка (например, english)"
// @Security     BearerAuth
// @Success      200  {object}  language.UserLanguageResponse  "ID записи user_languages"
// @Failure      400  {string}  string  "Invalid request (отсутствует параметр language, язык не найден в справочнике)"
// @Failure      401  {string}  string  "Unauthorized (отсутствует или недействителен токен)"
// @Failure      500  {string}  string  "Internal server error"
// @Router       /language [get]
func GetLanguage() {
	var _ = language.UserLanguageResponse{}
}
