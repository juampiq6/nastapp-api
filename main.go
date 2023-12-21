package main

import (
	"log"
	db "nastapp-api/db"
	r "nastapp-api/router"
	"os"

	"github.com/joho/godotenv"
)

// @title           Nastapp
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   Juampi Quinteros
// @contact.email  juampiq6@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	loadDotEnv()
	initMongoClient()
	runServer()
}

func runServer() {
	r := r.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
	defer db.KillMongoClient()
}

// Loads godotenv library to read env vars from .env file
func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initMongoClient() {
	db.InitMongoClient()
}
