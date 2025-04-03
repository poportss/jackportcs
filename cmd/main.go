package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/baseservice"
	"github.com/poportss/jackportcs/internal/database"
	"github.com/poportss/jackportcs/internal/pkg/auth"
	"github.com/poportss/jackportcs/internal/routes"
	"os"
)

func main() {

	db, err := database.ConnectDatabase()
	if err != nil {
		return
	}

	r := gin.Default()

	// Criar BaseService e injetar no serviço
	baseService := baseservice.NewBaseService(db)
	auth.AuthNewService(baseService)

	// 4️⃣ Configurar as rotas
	routes.SetupRoutes(r, db)

	apiPort := os.Getenv("API_PORT")
	// Se uma variável não estiver definida, exibe uma mensagem
	if apiPort == "" {
		fmt.Println("⚠️ API_PORT não definida! Usando padrão: 8080")
		apiPort = "8080"
	}

	fmt.Printf("🌍 API rodando na porta %s\n", apiPort)

	err = r.Run(":" + apiPort)
	if err != nil {
		return
	}
}
