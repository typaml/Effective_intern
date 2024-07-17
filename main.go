package main

import (
	"Effective_intern/config"
	"Effective_intern/controllers"
	_ "Effective_intern/docs" // Документация Swagger
	"Effective_intern/migrations"
	"Effective_intern/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// @title Time Tracker API
// @version 1.0
// @description This is a sample server for a time tracking application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	r := gin.Default()

	// Инициализация конфигурации и логгера
	config.InitConfig()
	utils.InitLogger()
	migrations.Migrate()
	// Маршрутизация
	userController := controllers.NewUserController()
	r.GET("/users", userController.GetUsers)
	r.GET("/users/:id", userController.GetUserByID)
	r.GET("/users/:id/tasks", userController.GetUserTasks)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
	r.POST("/users/tasks/:taskId/start", userController.StartTask)
	r.POST("/users/tasks/:taskId/stop", userController.StopTask)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
