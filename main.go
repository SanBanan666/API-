package main

import (
	"log"
	"os"

	"t3_juniorGo/database"
	"t3_juniorGo/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title People API
// @version 1.0
// @description API для работы с данными о людях
// @host localhost:8080
// @BasePath /
func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Printf("Ошибка загрузки .env файла: %v", err)
	}

	// Инициализация базы данных
	db := database.InitDB()

	// Инициализация роутера
	r := gin.Default()

	// Инициализация обработчиков
	personHandler := handlers.NewPersonHandler(db)

	// Маршруты API
	api := r.Group("/api")
	{
		people := api.Group("/people")
		{
			people.POST("", personHandler.CreatePerson)
			people.GET("", personHandler.GetPeople)
			people.PUT("/:id", personHandler.UpdatePerson)
			people.DELETE("/:id", personHandler.DeletePerson)
		}
	}

	// Swagger документация
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Получение порта из переменных окружения
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// Запуск сервера
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
