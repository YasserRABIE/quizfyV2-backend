package main

import (
	"log"
	"time"

	"github.com/YasserRABIE/QUIZFYv2/config"
	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/routes/questions"
	"github.com/YasserRABIE/QUIZFYv2/routes/quiz"
	"github.com/YasserRABIE/QUIZFYv2/routes/result"
	"github.com/YasserRABIE/QUIZFYv2/routes/session"
	"github.com/YasserRABIE/QUIZFYv2/routes/user"
	"github.com/gin-contrib/cors"
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
	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// assets
	r.Static("/uploads", "./uploads")

	// API group
	api_v1 := r.Group("/api/v1")
	// Register routes
	user.RegisterRoutes(api_v1)
	// Quiz routes
	quiz.QuizRoutes(api_v1)
	// Question routes
	questions.QuestionRoutes(api_v1)
	// Session routes
	session.SessionRoutes(api_v1)
	// Result routes
	result.ResultRoutes(api_v1)

	r.Run(config.GetEnv("PORT", ":3000"))
}
