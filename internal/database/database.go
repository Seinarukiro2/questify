package database

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User represents a user in the database.
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"` // User ID // ID почему то равен NULL
	Name      string    // User name
	CreatedAt time.Time // Creation timestamp
	Password  string    // User password
	Clubs     []Club    `gorm:"many2many:user_clubs;"` // User clubs
}

// Club represents a club in the database.
type Club struct {
	ID   uint   `gorm:"primaryKey"` // Club ID
	Name string `gorm:"unique"`     // Club name
}

// InitializeDatabase initializes the database connection and performs migrations.
func InitializeDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&User{}, &Club{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}
