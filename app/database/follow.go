package database

import "graphQL-go/app/models"

func IsAlreadyFollowing(followerID, followingID string) bool {
    var count int64
    DB.Model(&models.Follow{}).
        Where("follower_id = ? AND following_id = ?", followerID, followingID).
        Count(&count)
    return count > 0
}
