package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"src/internal/auth"
	configPkg "src/internal/config"
	"src/internal/db"
	"src/internal/middleware"
	"src/internal/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment")
	}

	// Загрузка порта из env
	port := os.Getenv("SERVER_PORT")

	// Загрузка конфигурации для jwt токенов
	jwt, err := configPkg.LoadJWTConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	//Подключение к бд
	database, err := db.Connect()
	if err != nil {
		log.Fatal("Could not connect to database: %w", err)

	}
	defer database.Close()

	// Подключение к сервису с email
	emailService, err := configPkg.ConnectSMTP()
	if err != nil {
		log.Fatal("Failed to connect to SMTP:", err)
	}

	// Конфигурация токенов
	authConfig := auth.Config{
		JWTSecretKey:    jwt.SecretKey,
		AccessTokenTTL:  jwt.AccessTokenTTL,
		RefreshTokenTTL: jwt.RefreshTokenTTL,
		VerificationTTL: jwt.VerificationTTL,
	}

	// Запуск обработчиков из пакета auth
	userStorage := auth.NewPostgresUserStorage(database)
	tsUserStorage := auth.NewPostgresTSUserStorage(database)
	refreshTokenStorage := auth.NewPostgresRefreshTokenStorage(database)
	verificationStorage := auth.NewMemoryVerificationStorage()
	resetPasswordStorage := auth.NewMemoryResetPasswordStorage()

	authService := auth.NewAuthManager(userStorage, refreshTokenStorage, verificationStorage, resetPasswordStorage, tsUserStorage, emailService, authConfig)
	authHandler := auth.NewHandler(authService)

	authMiddleware := middleware.NewAuthMiddleware(jwt.SecretKey)

	//Пути - src/internal/router/router.go
	router := router.New(authMiddleware, authHandler)

	// Запуск сервера
	log.Printf("Сервер запущен на http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
