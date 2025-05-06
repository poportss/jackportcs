package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/baseservice"
	"github.com/poportss/jackportcs/internal/database"
	"github.com/poportss/jackportcs/internal/routes"
	"os"
)

func main() {
	//testando o bui;d
	db, err := database.ConnectDatabase()
	if err != nil {
		return
	}

	r := gin.Default()

	// Configuração do CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // URL do seu frontend
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	r.Use(cors.New(config))

	baseService := baseservice.NewBaseService(db)

	routes.SetupRoutes(r, baseService)

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		fmt.Println("⚠️ API_PORT não definida! Usando padrão: 8080")
		apiPort = "8080"
	}

	fmt.Printf("🌍 API rodando na porta %s\n", apiPort)

	// 🔎 Lista todos os endpoints
	for _, ri := range r.Routes() {
		fmt.Printf("🔗 %s %s\n", ri.Method, ri.Path)
	}

	err = r.Run(":" + apiPort)
	if err != nil {
		return
	}
}
