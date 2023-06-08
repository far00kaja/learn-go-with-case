package main

import (
	"os"

	_ "github.com/far00kaja/learn-go-with-case/docs"
	"github.com/far00kaja/learn-go-with-case/server"
	"github.com/joho/godotenv"
)

// @title Tag Auth Service API
// @Version 1.0
// @description A Auth Service API in Go using Gin Framework

// @host localhost:9997
//  @basepath /
func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	r := server.Init()
	r.Run("0.0.0.0:" + PORT)

}
