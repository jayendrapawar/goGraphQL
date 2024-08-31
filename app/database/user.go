package database

import "graphQL-go/app/models"

func GetUserByID(userID string) (models.User, error) {
    var user models.User
    result := DB.First(&user, "id = ?", userID)
    return user, result.Error
}

func UserExists(userID string) bool {
    var user models.User
    result := DB.First(&user, "id = ?", userID)
    return result.RowsAffected > 0
}