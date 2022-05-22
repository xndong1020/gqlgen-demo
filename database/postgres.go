package database

import (
	"fmt"

	"acy.com/gqlgendemo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
    Host     string
    Port     string
    Password string
    User     string
    DBName   string
    SSLMode  string
    Schema   string
}

func NewConnection(config *Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
        config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode, config.Schema,
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return db, err
    }
    return db, nil
}

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&models.Book{})
}