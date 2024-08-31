package main

import (
	"graphQL-go/service/database"
	"graphQL-go/service/graphql"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
    // Database connection
    _, err := database.Connect()
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    schema := graphql.NewSchema()

    h := handler.New(&handler.Config{
        Schema:   &schema,
        Pretty:   true,
        GraphiQL: true,
    })

    http.Handle("/graphql", h)

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
