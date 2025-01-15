package service

import (
	"Notes/pkg/jwt_token"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

const secretKey = "khetag_pig"

// CreateUser - регистрация пользователя
func (s *Service) CreateUser(c echo.Context) error {
	newUser := new(RegisterUser)
	if err := c.Bind(newUser); err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusBadRequest, "Некорректные данные")
	}

	if newUser.Name == "" || newUser.Password == "" || newUser.Email == "" {
		return c.JSON(http.StatusBadRequest, "Все поля обязательны для заполнения")
	}

	id := strconv.FormatInt(time.Now().UnixNano(), 10)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		s.logger.Error("Ошибка хеширования пароля:", err)
		return c.JSON(http.StatusInternalServerError, "Ошибка при обработке пароля")
	}

	fmt.Println(id, "ИД пользователя")
	token, err := jwt_token.CreateToken(id, newUser.Name, secretKey)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusBadRequest, "Некорректные данные")
	}

	repo := s.registerUserRepo
	err = repo.CreateNewUser(id, newUser.Name, newUser.Email, string(hashedPassword), token)
	if err != nil {
		s.logger.Error("Ошибка сохранения пользователя:", err)
		return c.JSON(http.StatusConflict, "Пользователь уже существует")
	}

	return c.JSON(http.StatusOK, "Вы успешно зарегистрированы")
}

// AuthUser - авторизация пользователя
func (s *Service) AuthUser(c echo.Context) error {
	user := new(AuthUser)
	if err := c.Bind(user); err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusBadRequest, "Некорректные данные")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		s.logger.Error("Ошибка хеширования пароля:", err)
		return c.JSON(http.StatusInternalServerError, "Ошибка при обработке пароля")
	}

	repo := s.registerUserRepo
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
		fmt.Println(token)
		if len(token) == 0 {
			return c.JSON(http.StatusBadRequest, "Отсутствует токен")
		}

		jwt, err := jwt_token.DecodeJWT(token, []byte(secretKey))
		fmt.Println(jwt["sub"])
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, "Не удалось декодировать токен")
		}

		return apiFunc(c, jwt["sub"].(string))
	}
}
