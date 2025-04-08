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
	maxConn, err := strconv.Atoi(maxConnStr)
	if err != nil || maxConn <= 0 {
		maxConn = 10
	}

	base.SetMaxIdleConns(maxConn)
	base.SetMaxOpenConns(maxConn / 2)

	if err := base.Ping(); err != nil {
		log.Fatal("❌ Erro ao conectar no banco:", err)
	}
	log.Println("✅ Banco de dados conectado e respondendo!")

	return db, nil
}
