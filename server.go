package main

import (
	"golanggraphql/config"
	"golanggraphql/db"
	"golanggraphql/graph"
	"golanggraphql/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("app.env")
	config.Initconfig()
	db.InitDatabase()
	router := chi.NewRouter()
	router.Handle("/"+config.GetConfig().APPVersion+"/query", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	http.ListenAndServe(":"+config.GetConfig().Port, router)
}
