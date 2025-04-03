package database

import (
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDatabase estabelece conexão com o banco de dados
func ConnectDatabase() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DB_CONN"),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
		Logger:                                   newLogger,
	})
	if err != nil {
		return nil, err
	}
	base, err := db.DB()
	if err != nil {
		return nil, err
	}

	maxConnStr := os.Getenv("DB_MAX_CONN")
	// Definir um valor padrão para maxConn se estiver vazio
	maxConn, err := strconv.Atoi(maxConnStr)
	if err != nil || maxConn <= 0 {
		maxConn = 10 // Valor padrão caso não seja possível converter
	}

	base.SetMaxIdleConns(maxConn)
	base.SetMaxOpenConns(maxConn / 2)
	dsn := base.Stats() // Isso não retorna o nome do banco, mas confirma a conexão ativa
	log.Println("✅ Conectado ao banco de dados com sucesso!")
	log.Println(dsn)
	return db, nil
}
