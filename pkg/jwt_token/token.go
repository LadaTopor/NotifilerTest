package jwt_token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func CreateToken(id, name, key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("JWT_SECRET_KEY is not set")
	}
	payload := jwt.MapClaims{
		"sub":  id,
		"name": name,
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString([]byte(key))
	if err != nil {
		return "", errors.New("ошибка создания JWT токена")
	}
	return t, nil
}

func DecodeJWT(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	// Парсинг и верификация токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("ошибка при разборе токена: %w", err)
	}

	// Извлекаем claims из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("не удалось преобразовать claims")
	}

	return claims, nil
}
