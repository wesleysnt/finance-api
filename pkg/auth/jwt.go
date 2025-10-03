package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wesleysnt/finance-api/app/config"
)

type JWTService interface {
	GenerateToken(userId int, email, role string) (string, error)
	GenerateRefreshToken(userId int) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type jwtService struct {
	secretKey     string
	issuer        string
	tokenExpiry   time.Duration
	refreshExpiry time.Duration
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey:     config.Conf.Jwt.Secret,
		tokenExpiry:   time.Duration(config.Conf.Jwt.Expiry) * time.Hour,
		refreshExpiry: time.Duration(config.Conf.Jwt.Refresh) * time.Hour,
	}
}
func NewJWTServiceWithConfig(secretKey, issuer string, tokenExpiry, refreshExpiry time.Duration) JWTService {
	return &jwtService{
		secretKey:     secretKey,
		issuer:        issuer,
		tokenExpiry:   tokenExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (j *jwtService) GenerateToken(userId int, email, role string) (string, error) {
	claims := &Claims{
		UserID: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) GenerateRefreshToken(userId int) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshExpiry)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.secretKey)
}

func (j *jwtService) ValidateToken(tokenString string) (*Claims, error) {
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
