package graphql

import (
    "github.com/graphql-go/graphql"
)

var followMutation = graphql.NewObject(graphql.ObjectConfig{
    Name: "Mutation",
    Fields: graphql.Fields{
        "follow": &graphql.Field{
            Type: graphql.String,
            Args: graphql.FieldConfigArgument{
                "follower_id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
                "following_id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
            },
            Resolve: followResolver,
        },
        "unfollow": &graphql.Field{
            Type: graphql.String,
            Args: graphql.FieldConfigArgument{
                "follower_id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
                "following_id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
            },
            Resolve: unfollowResolver,
        },
    },
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        "followers": &graphql.Field{
            Type: graphql.NewList(userType),
            Args: graphql.FieldConfigArgument{
                "user_id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
            },
            Resolve: listFollowersResolver,
        },
        "following": &graphql.Field{
            Type: graphql.NewList(userType),
            Args: graphql.FieldConfigArgument{
                "user_id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
            },
            Resolve: listFollowingResolver,
        },
    },
})

func NewSchema() graphql.Schema {
    schema, _ := graphql.NewSchema(
        graphql.SchemaConfig{
            Query:    queryType,
            Mutation: followMutation,
        },
    )
    return schema
}
