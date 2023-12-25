package database

import (
	"fmt"
)

// CreateUser creates a new user in the database.
func CreateUser(user *User) error {
	db, err := InitializeDatabase()
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	result := db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("failed to create user: %w", result.Error)
	}
	return nil
}
