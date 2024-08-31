package database

import "graphQL-go/app/models"

func IsAlreadyFollowing(followerID, followingID string) bool {
    var count int64
    DB.Model(&models.Follow{}).
        Where("follower_id = ? AND following_id = ?", followerID, followingID).
        Count(&count)
    return count > 0
}

func ListFollowers(userID string) ([]models.User, error) {
    var users []models.User
    DB.Table("users").
        Joins("JOIN follows ON users.id = follows.follower_id").
        Where("follows.following_id = ?", userID).
        Scan(&users)
    return users, nil
}

func ListFollowing(userID string) ([]models.User, error) {
    var users []models.User
    DB.Table("users").
        Joins("JOIN follows ON users.id = follows.following_id").
        Where("follows.follower_id = ?", userID).
        Scan(&users)
    return users, nil
}
