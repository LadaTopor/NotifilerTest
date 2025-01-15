package service

import (
	"Notes/internal/notes"
	"Notes/internal/users"
	"database/sql"

	"github.com/labstack/echo/v4"
)

const (
	InvalidParams       = "invalid params"
	InternalServerError = "internal error"
	InvalidBody         = "Invalid request body"
)

type Service struct {
	db     *sql.DB
	logger echo.Logger

	notesRepo        *notes.Repo
	registerUserRepo *users.Repo
}

func NewService(db *sql.DB, logger echo.Logger) *Service {
	svc := &Service{
		db:     db,
		logger: logger,
	}
	svc.initRepositories(db)

	return svc
}

func (s *Service) initRepositories(db *sql.DB) {
	s.notesRepo = notes.NewRepo(db)
	s.registerUserRepo = users.NewRepo(db)
}

// Пока можно не вдаваться в то что ниже

type Response struct {
	Object       any    `json:"object,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

func (r *Response) Error() string {
	return r.ErrorMessage
}

func (s *Service) NewError(err string) (int, *Response) {
	return 400, &Response{ErrorMessage: err}
}
