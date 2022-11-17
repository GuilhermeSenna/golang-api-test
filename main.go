package main

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/guilhermeSenna/golang-api-test/config"
	"github.com/guilhermeSenna/golang-api-test/routes"
)

func loadEnvPort() (string) {
	err := godotenv.Load()

	if (err != nil || os.Getenv("API_PORT") == "") {
		fmt.Printf("\033[1;33m%s\033[0m", "Error loading .env file or required variable!\n")
		return "8080"
	}else{
		return os.Getenv("API_PORT")
	}
}

func main() {
	port := loadEnvPort()
	router := gin.New()
	config.Connect()
	routes.Routes(router)
	router.Run("127.0.0.1:" + port)
}