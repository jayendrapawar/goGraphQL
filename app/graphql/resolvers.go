package graphql

import (
	"graphQL-go/app/services"

	"github.com/graphql-go/graphql"
)

func followResolver(p graphql.ResolveParams) (interface{}, error) {
    followerID := p.Args["follower_id"].(string)
    followingID := p.Args["following_id"].(string)
    return nil, services.FollowUser(followerID, followingID)
}

func unfollowResolver(p graphql.ResolveParams) (interface{}, error) {
    followerID := p.Args["follower_id"].(string)
    followingID := p.Args["following_id"].(string)
    return nil, services.UnfollowUser(followerID, followingID)
}

func listFollowersResolver(p graphql.ResolveParams) (interface{}, error) {
    userID := p.Args["user_id"].(string)
    return services.ListFollowers(userID)
}

func listFollowingResolver(p graphql.ResolveParams) (interface{}, error) {
    userID := p.Args["user_id"].(string)
    return services.ListFollowing(userID)
}
