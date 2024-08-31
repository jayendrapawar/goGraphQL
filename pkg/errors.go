package errors

import "errors"

var (
    ErrUserNotFound    = errors.New("user not found")
    ErrFollowExists    = errors.New("already following this user")
    ErrCannotFollowSelf = errors.New("cannot follow yourself")
)
