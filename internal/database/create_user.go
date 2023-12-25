package database

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("failed to create user: %w", result.Error)
	}
	return nil
}
