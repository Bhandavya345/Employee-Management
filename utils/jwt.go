package utils

import (
	"errors"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("mysecretkey")

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	RoleID uint   `json:"roleID"`
	jwt.RegisteredClaims
}

// Generate JWT Token
func GenerateJWT(userID uint, email, role string, roleID uint) (string, error) {

	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RoleID: roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SecretKey)
}

// Validate JWT Token
func ValidateJWT(tokenString string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
