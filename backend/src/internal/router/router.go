package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "src/docs"
	"src/internal/auth"
	"src/internal/language"
	"src/internal/middleware"
)

func New(authMiddleware *middleware.AuthMiddleware, authHandler *auth.Handler, langHandler *language.Handler) http.Handler {
	r := chi.NewRouter()

	// Глобальные middleware для всех запросов
	r.Use(chimiddleware.Logger)    // логирование запросов
	r.Use(chimiddleware.Recoverer) // восстановление после паник

	// Swagger UI — доступен по /swagger/index.html
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // обязательно!
	))

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/verify", authHandler.VerifyCode)
		r.Post("/login", authHandler.Login)
		r.Post("/refresh", authHandler.Refresh)
		r.Post("/logout", authHandler.Logout)
		r.Post("/forgot-password", authHandler.ForgotPassword)
		r.Post("/verify-reset-code", authHandler.VerifyResetCode)
		r.Post("/set-password", authHandler.SetPassword)
	})

	r.Route("/language", func(r chi.Router) {
		// Защищённый маршрут
		r.With(authMiddleware.Authenticate).Get("/", langHandler.GetOrCreateUserLanguage)
	})

	return r
}
