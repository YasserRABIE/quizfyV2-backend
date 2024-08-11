package main

import (
	"log"

	"github.com/YasserRABIE/QUIZFYv2/config"
	"github.com/YasserRABIE/QUIZFYv2/db"
	routes "github.com/YasserRABIE/QUIZFYv2/routes/user"
	"github.com/gin-gonic/gin"
)

func init() {
	config.InitConfig()

	if err := db.InitDB(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	if err := db.InitTables(); err != nil {
		log.Fatalf("Table migration failed: %v", err)
	}
}

func main() {
	r := gin.Default()

	api_v1 := r.Group("/api/v1")
	// Register routes
	routes.RegisterRoutes(api_v1)

	r.Run(config.GetEnv("PORT", ":3000"))
}
