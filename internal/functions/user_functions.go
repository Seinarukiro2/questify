package functions

import (
	"errors"
	"fmt"
	"questify/internal/database"
)

func CreateUserAndPrint(name, password string) error {
	newUser := &database.User{
		Name:     name,
		Password: password,
	}

	err := database.CreateUser(newUser)
	if err != nil {
		return errors.New(err.Error())
	}

	fmt.Printf("New user was created: %s\n", name)
	return nil
}
