package util

// import (
// 	"assignment-golang-backend/dto"
// 	"assignment-golang-backend/entity"
// 	"assignment-golang-backend/httperror"
// 	"os"
// 	"time"

// 	"github.com/golang-jwt/jwt/v4"
// 	"golang.org/x/crypto/bcrypt"
// )

// func GenerateAccessToken(user *entity.User) (*dto.TokenResponse, error) {

// 	claims := jwt.MapClaims{
// 		"id":  user.UserId,
// 		"iss": "localhost:8000/",
// 		"iat": time.Now().Unix(),
// 		"exp": time.Now().Add(time.Hour * 24).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	hMacSecret := []byte(os.Getenv("SECRET_KEY"))
// 	tokenString, err := token.SignedString(hMacSecret)
// 	if err != nil {
// 		return nil, httperror.ErrFailedCreateToken
// 	}

// 	return &dto.TokenResponse{Token: tokenString}, nil
// }

// func ComparePassword(hashedPwd string, inputPwd string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
// 	return err == nil
// }

// func HashPassword(password string) (string, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
// 	if err != nil {
// 		return "", httperror.ErrGenerateHash
// 	}
// 	return string(hashedPassword), nil
// }
