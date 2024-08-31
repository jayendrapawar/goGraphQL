package services

import "graphQL-go/app/database"

func GetUserByID(userID string) (interface{}, error) {
    return database.GetUserByID(userID)
}
