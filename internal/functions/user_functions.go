package functions

import (
    "fmt"
    "yourproject/database"
)

func CreateUserAndPrint(db *gorm.DB, name, password string) error {
    newUser := &database.User{
        Name:     name,
        Password: password,
    }

    err := database.CreateUser(db, newUser)
    if err != nil {
        return err
    }

    fmt.Printf("New user was created: %s\n", name)
    return nil
}