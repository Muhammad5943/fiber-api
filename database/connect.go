package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Muhammad5943/fiber-api/config"
	"github.com/Muhammad5943/fiber-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Declear variable for database
// var DB *gorm.DB

type DbInstance struct {
	Db *gorm.DB
}

// Declear variable for database
var Database DbInstance

// connecDB to connect DB
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("Database Not Connected")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	// Connection url to connect to postgres database
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connected to Opened Database")
	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	fmt.Println("Database Migrated")

	Database = DbInstance{
		Db: db,
	}
}
