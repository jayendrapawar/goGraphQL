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

    },
})


func NewSchema() graphql.Schema {
    schema, _ := graphql.NewSchema(
        graphql.SchemaConfig{
            Mutation: followMutation,
        },
    )
    return schema
}
