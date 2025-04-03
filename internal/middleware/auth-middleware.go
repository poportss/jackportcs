package middleware

import (
	"github.com/poportss/jackportcs/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware protege as rotas autenticadas
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		// Validar o token usando a função do utils/token.go
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Adiciona os dados do usuário no contexto para serem usados em outras rotas
		c.Set("username", claims["username"])
		c.Next()
	}
}
