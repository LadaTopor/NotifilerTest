package main

import (
	"Notes/internal/service"
	"Notes/pkg/logs"

	"github.com/labstack/echo/v4"
)

func main() {
	// создаем логгер
	logger := logs.NewLogger(false)

	// подключаемся к базе
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger)

	router := echo.New()
	// создаем группу api
	api := router.Group("api")

	// прописываем пути
	api.GET("/notes", svc.CheckAuth(svc.GetNotes))
	api.GET("/note/:id", svc.CheckAuth(svc.GetNoteById))
	api.POST("/note", svc.CheckAuth(svc.CreateNote))
	api.PUT("/note", svc.CheckAuth(svc.UpdateNoteById))
	api.DELETE("/note", svc.CheckAuth(svc.DeleteNoteById))

	api.POST("/register", svc.CreateUser)
	api.POST("/auth", svc.AuthUser)

	// запускаем сервер, чтобы слушал 8000 порт1
	router.Logger.Fatal(router.Start(":8000"))
}
