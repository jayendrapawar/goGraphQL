package services

import (
	"errors"
	"graphQL-go/app/database"
	"graphQL-go/app/models"
)

func FollowUser(followerID, followingID string) error {

    if !database.UserExists(followerID) {
        return errors.New("follower user does not exist")
    }
    if !database.UserExists(followingID) {
        return errors.New("following user does not exist")
    }
    
    if followerID == followingID {
        return errors.New("cannot follow yourself")
    }

    if database.IsAlreadyFollowing(followerID, followingID) {
        return errors.New("already following this user")
    }

    follow := models.Follow{
        FollowerID:  followerID,
        FollowingID: followingID,
    }

    return database.DB.Create(&follow).Error
}

func UnfollowUser(followerID, followingID string) error {
    return database.DB.Where("follower_id = ? AND following_id = ?", followerID, followingID).Delete(&models.Follow{}).Error
}

func ListFollowers(userID string) ([]models.User, error) {
    return database.ListFollowers(userID)
}

func ListFollowing(userID string) ([]models.User, error) {
    return database.ListFollowing(userID)
}
