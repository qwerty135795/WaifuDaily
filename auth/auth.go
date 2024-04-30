package auth

import (
	"Chat/data"
	"Chat/dto"
	"Chat/models"
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"log"
)

func Register(dto dto.RegisterDto) (int, error) {
	saltSize := 12
	salt := generateSalt(saltSize)
	pass := generatePassword(dto.Password, salt)
	user := models.User{
		Name:         dto.Name,
		Gender:       dto.Gender,
		Age:          dto.Age,
		Waifu:        dto.Waifu,
		Password:     pass,
		PasswordSalt: salt,
	}
	id, err := data.CreateUser(user)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return id, nil
}

func Login(name, password string) (string, error) {
	user, err := data.GetUserByName(name)
	if err != nil {
		return "", err
	}
	if checkPassword(user, password) {
		token, err := GenerateToken(*user)
		if err != nil {
			return "", nil
		}
		return token, nil
	}
	return "", errors.New("NotFound")
}

func checkPassword(user *models.User, password string) bool {
	hash := generatePassword(password, user.PasswordSalt)
	return bytes.Equal(user.Password, hash)
}

func generatePassword(password string, salt []byte) []byte {
	pass := []byte(password)
	sha512Hasher := sha512.New()

	pass = append(pass, salt...)

	sha512Hasher.Write(pass)
	return sha512Hasher.Sum(nil)

}

func generateSalt(n int) []byte {
	salt := make([]byte, n)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal(err)
	}
	return salt
}
