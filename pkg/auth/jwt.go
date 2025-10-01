package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wesleysnt/finance-api/app/config"
)

type JWTService struct {
	secretKey     string
	issuer        string
	tokenExpiry   time.Duration
	refreshExpiry time.Duration
}

func NewJWTService() *JWTService {
	return &JWTService{
		secretKey:     config.Conf.Jwt.Secret,
		tokenExpiry:   time.Duration(config.Conf.Jwt.Expiry) * time.Hour,
		refreshExpiry: time.Duration(config.Conf.Jwt.Refresh) * time.Hour,
	}
}

func (j *JWTService) GenerateToken(userId int, email, role string) (string, error) {
	claims := &Claims{
		UserID: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    j.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) GenerateRefreshToken(userId int) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshExpiry)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    j.issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.secretKey)
}

func (j *JWTService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
