package models

type Follow struct {
    ID          uint   `gorm:"primaryKey"`
    FollowerID  string `gorm:"not null"`
    FollowingID string `gorm:"not null"`
}
