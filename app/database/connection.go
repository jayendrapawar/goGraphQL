package database

import (
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
    dsn := "root:1234567890@tcp(localhost:3306)/graphql?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    log.Println("MySQL database connected successfully")
    DB = db
    return db, nil
}
