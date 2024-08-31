package services

import (
	"errors"
	"graphQL-go/app/database"
	"graphQL-go/app/models"
)

func FollowUser(followerID, followingID string) error {
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