package main

import (
	orm "graphql-go-template/internal/database"
	"graphql-go-template/internal/handlers"
	log "log"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"graphql-go-template/config"
)

var (
	env    config.EnvConfig
	logger *zap.Logger
)

var gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	gqlPath = "/graphql"
	gqlPgPath = "/"
	isPgEnabled = true
}

func loadEnv() {
	if err := config.Process(&env); err != nil {
		log.Fatal("load config error > ", err)
	}
}

// Run spins up the server
func main() {
	loadEnv()

	// log.Println("GORM_CONNECTION_DSN: ", utils.MustGet("GORM_CONNECTION_DSN"))
	host := "localhost"

	// Create a new ORM instance to send it to our server
	orm, err := orm.Factory(env.Database)
	if err != nil {
		log.Panic(err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// AllowOriginFunc:  func(origin string) bool { return origin == "http://localhost:3000" },
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Handlers
	// Simple keep-alive/ping handler
	r.GET("/heartbeat", handlers.Heartbeat())

	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + strconv.Itoa(env.Env.HTTPPort) + gqlPgPath)
	}

	// Pass in the ORM instance to the GraphqlHandler
	r.POST(gqlPath, handlers.GraphqlHandler(orm))
	log.Println("GraphQL @ " + strconv.Itoa(env.Env.HTTPPort))

	// Run the server
	// Inform the user where the server is listening
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatal(r.Run(host + ":" + strconv.Itoa(env.Env.HTTPPort)))
}
