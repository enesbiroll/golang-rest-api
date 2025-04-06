package jwt

import (
	"errors"
	"time"

	"rest-api/models"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key") // Secret key for JWT

// GenerateJWT token oluşturma
func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":      user.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token 1 gün geçerli olacak
		"iat":      time.Now().Unix(),
		"username": user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ParseJWT token doğrulama
func ParseJWT(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("could not parse claims")
	}

	userID := uint(claims["sub"].(float64))
	username := claims["username"].(string)

	// Token'dan gelen kullanıcı bilgilerini döndür
	return &models.User{
		ID:       userID,
		Username: username,
	}, nil
}
