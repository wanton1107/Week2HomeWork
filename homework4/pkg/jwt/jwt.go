package jwt

import (
	"fmt"
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	jwtv5.RegisteredClaims
}

type Manager struct {
	secret []byte
	issuer string
	expire time.Duration
}

func NewManager(secret, issuer string, expireHour int) *Manager {
	return &Manager{
		secret: []byte(secret),
		issuer: issuer,
		expire: time.Duration(expireHour) * time.Hour,
	}
}

func (m *Manager) GenerateToken(userID uint, username string) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwtv5.RegisteredClaims{
			Issuer:    m.issuer,
			ExpiresAt: jwtv5.NewNumericDate(now.Add(m.expire)),
			IssuedAt:  jwtv5.NewNumericDate(now),
		},
	}

	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}

func (m *Manager) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwtv5.ParseWithClaims(tokenString, &Claims{}, func(token *jwtv5.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtv5.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.secret, nil
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
