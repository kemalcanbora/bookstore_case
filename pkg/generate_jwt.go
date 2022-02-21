package pkg

import (
	"bookstore_case/models"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

func GenerateJWT(user models.User) (string, error) {
	GetEnv()
	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = user.Email
	claims["first_name"] = user.FirstName
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("Something Went Wrong: %s\n", err.Error())
		return "", err
	}
	return tokenString, nil
}
