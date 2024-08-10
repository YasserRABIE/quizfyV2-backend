package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quizfy/api/config"
	"github.com/quizfy/api/db"
)

func init() {
	config.InitConfig()
	db.InitDB()
	db.InitTables()
}

func main() {
	r := gin.Default()

	r.Run(config.GetEnv("PORT", ":3000"))
}
