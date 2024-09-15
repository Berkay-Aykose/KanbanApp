package controllers

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTKey = []byte("17628D055CAA43C6D4E658A40EC26273D25681934E2582D52A61CEFF226DE4A5") //secret key

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token geçerlilik süresi 24 saat olarak ayarlandı.
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func VerifyToken(tokenString string) (*Claims, error) {
	token1, err := ExtractTokenFromHeader(tokenString)
	if err != nil {
		return nil, err
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(token1, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return JWTKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ExtractTokenFromHeader(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

func ExtractUserIDAndExpirationFromToken(tokenString string) (string, int64, error) {
	claims, err := VerifyToken(tokenString)
	if err != nil {
		return "", 0, err
	}

	return claims.UserID, claims.ExpiresAt, nil
}
