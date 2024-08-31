package database

import "graphQL-go/app/models"

func GetUserByID(userID string) (models.User, error) {
    var user models.User
    result := DB.First(&user, "id = ?", userID)
    return user, result.Error
}
