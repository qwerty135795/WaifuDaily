package auth

import (
	"Chat/models"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var apiSecret = []byte("NewSeCrEtStRiNgSSsIl0VeeeY000U")

func GenerateToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["id"] = user.Id
	claims["username"] = user.Name
	s, err := token.SignedString(apiSecret)
	if err != nil {
		return "", err
	}
	return s, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return apiSecret, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("Token invalid")
	}
	return nil
}
