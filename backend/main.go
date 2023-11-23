package main

import (
	"task-canvas/auth"
	"task-canvas/interface/graph"
	"task-canvas/interface/graph/generated"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/patrickmn/go-cache"

	"github.com/gin-gonic/gin"
)

// Note: 参考 https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/gin.md
func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func main() {
	r := gin.Default()
	c := cache.New(time.Hour, 2*time.Hour)

	r.Use(auth.Middleware(c))

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"HEAD",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Authorization",
			"Baggage",
		},
		AllowCredentials: true,
	}))

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
