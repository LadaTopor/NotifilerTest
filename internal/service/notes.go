package service

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// CreateNote - создание заметки
func (s *Service) CreateNote(c echo.Context, userId string) error {
	note := new(Note)
	err := c.Bind(&note)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	resp, err := http.Get("https://favqs.com/api/qotd")
	if err != nil {
		s.logger.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.logger.Error("Неудачный запрос, статус:", resp.StatusCode)
	}

	var jsonResponse QuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		s.logger.Error(err)
	}

	quoteBody := jsonResponse.Quote.Body

	repo := s.notesRepo
	body := fmt.Sprintf("%s %s", note.Body, quoteBody)
	err = repo.CreateNewNote(note.Title, body, userId)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, "OK")
}

// GetNotes - все запросы
func (s *Service) GetNotes(c echo.Context, userId string) error {
	repo := s.notesRepo
	words, err := repo.GetAllNotes(userId, 0)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, words)
}

// GetNoteById - Достать заметку по id
func (s *Service) GetNoteById(c echo.Context, userId string) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.notesRepo
	words, err := repo.GetAllNotes(userId, id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, words)
}

// DeleteNoteById - удаление заметки
func (s *Service) DeleteNoteById(c echo.Context, userId string) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.notesRepo
	err = repo.DeleteNote(id, userId)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, "OK")
}

// UpdateNoteById - изменение заметки
func (s *Service) UpdateNoteById(c echo.Context, userId string) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	note := new(Note)
	err = c.Bind(&note)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.notesRepo
	err = repo.UpdateNote(id, userId, note.Title, note.Body)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, "OK")
}
