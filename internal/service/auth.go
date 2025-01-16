package service

import (
	"Notes/pkg/jwt_token"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

const secretKey = "khetag_pig"

// CreateUser - регистрация пользователя
func (s *Service) CreateUser(c echo.Context) error {
	newUser := new(User)
	if err := c.Bind(newUser); err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusBadRequest, "Некорректные данные")
	}

	if len(newUser.Name) == 0 || len(newUser.Password) == 0 || len(newUser.Email) == 0 {
		return c.JSON(http.StatusBadRequest, "Все поля обязательны для заполнения")
	}

	id := strconv.FormatInt(time.Now().UnixNano(), 10)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		s.logger.Error("Ошибка хеширования пароля:", err)
		return c.JSON(http.StatusInternalServerError, "Ошибка при обработке пароля")
	}

	token, err := jwt_token.CreateToken(id, newUser.Name, secretKey)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusBadRequest, "Некорректные данные")
	}

	repo := s.UserRepo
	err = repo.CreateNewUser(id, newUser.Name, newUser.Email, string(hashedPassword), token)
	if err != nil {
		s.logger.Error("Ошибка сохранения пользователя:", err)
		return c.JSON(http.StatusConflict, "Пользователь уже существует")
	}

	return c.JSON(http.StatusOK, "Вы успешно зарегистрированы")
}

// AuthUser - авторизация пользователя
func (s *Service) AuthUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusBadRequest, "Некорректные данные")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		s.logger.Error("Ошибка хеширования пароля:", err)
		return c.JSON(http.StatusInternalServerError, "Ошибка при обработке пароля")
	}

	repo := s.UserRepo
	err = repo.VerifyingUserData(user.Email, string(hashedPassword))
	if err != nil {
		s.logger.Error("Ошибка введённых данных:", err)
		return c.JSON(http.StatusConflict, "Неверный пароль или email")
	}

	return c.JSON(http.StatusOK, "Вы успешно авторизированны")
}

func (s *Service) CheckAuth(apiFunc func(c echo.Context, userId string) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if len(token) == 0 {
			return c.JSON(http.StatusBadRequest, "Отсутствует токен")
		}

		jwt, err := jwt_token.DecodeJWT(token, []byte(secretKey))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Не удалось декодировать токен")
		}

		return apiFunc(c, jwt["sub"].(string))
	}
}
