package graphql

import (
    "github.com/graphql-go/graphql"
)

func NewSchema() graphql.Schema {
    schema, _ := graphql.NewSchema(
        graphql.SchemaConfig{
        },
    )
    return schema
}
