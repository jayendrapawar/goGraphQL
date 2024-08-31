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

