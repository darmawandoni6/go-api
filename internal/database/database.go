package database

import (
	"fmt"
	"log"

	"github.com/donidarmawan/test-go/internal/config"
	"github.com/donidarmawan/test-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// AllModels daftar semua model untuk AutoMigrate
func AllModels() []interface{} {
	return []interface{}{
		&models.User{},
		&models.Product{},
		&models.Transaction{},
		&models.TransactionItem{},
		&models.Debt{},
		&models.DebtPayment{},
		&models.CashFlow{},
		&models.Category{},
		&models.StockMovement{},
	}
}

// Connect opens a GORM connection to PostgreSQL and runs auto-migration.
func Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")

	if err := db.AutoMigrate(AllModels()...); err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}

	// seed category
	initialCategories := []string{"Makanan", "Minuman", "Sembako", "Rokok", "Lainnya"}
	for _, name := range initialCategories {
		db.Where(models.Category{Name: name}).
			FirstOrCreate(&models.Category{Name: name})
	}

	log.Println("Database migration completed")
	return db
}
