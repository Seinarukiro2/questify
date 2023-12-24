package store

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)


func SetupDatabase() (*gorm.DB, error) {
    dsn := "databaseurl"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}
