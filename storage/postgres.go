package storage

import (
	"fmt"
	"go-shop-diplom/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %w", err)
	}

	// Настройка пула соединений
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}

	// Устанавливаем максимальное количество открытых соединений в пуле
	sqlDB.SetMaxOpenConns(10)

	// Устанавливаем максимальное количество простаивающих соединений
	sqlDB.SetMaxIdleConns(10)

	// Устанавливаем максимальное время жизни соединения
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Устанавливаем максимальное время простоя соединения
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)

	log.Printf("Successfully connected to database: %s with connection pool (max connections: 10)", os.Getenv("DB_NAME"))
}

func MigrateModels() {
	models.MigrateProduct(DB)
	models.MigrateUser(DB)
	models.MigrateBasket(DB)
}
