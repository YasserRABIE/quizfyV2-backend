package main

import (
	"github.com/quizfy/api/config"
	"github.com/quizfy/api/db"
)

func init() {
	config.InitConfig()
	db.InitDB()
}

func main() {

}
