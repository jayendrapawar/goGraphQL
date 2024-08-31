package main

import (
	"graphQL-go/app/database"
	"graphQL-go/app/graphql"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
    // Initialize the database connection
    _, err := database.Connect()
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    // Setup GraphQL schema
    schema := graphql.NewSchema()

    // Create a new GraphQL HTTP handler
    h := handler.New(&handler.Config{
        Schema:   &schema,
        Pretty:   true,
        GraphiQL: true,
    })

    // Serve the GraphQL endpoint
    http.Handle("/graphql", h)

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
