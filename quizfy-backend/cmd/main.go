package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/quizfy/api/config"
	"github.com/quizfy/api/db"
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

	r.Run(config.GetEnv("PORT", ":3000"))
}
