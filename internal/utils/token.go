package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Chave secreta para assinar o token (NÃO EXIBA ISSO EM PRODUÇÃO)
var jwtSecret = []byte("minha_chave_secreta")

// GenerateToken cria um token JWT válido por 1 hora
func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Expiração em 1 hora
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken verifica a autenticidade do token e retorna as claims
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("assinatura inválida")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token inválido")
	}

	// Retorna as claims do token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("não foi possível obter as claims")
	}

	return claims, nil
}
