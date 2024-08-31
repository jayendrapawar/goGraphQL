package errors

import "errors"

var (
    ErrFollowExists    = errors.New("already following this user")
    ErrCannotFollowSelf = errors.New("cannot follow yourself")
)
