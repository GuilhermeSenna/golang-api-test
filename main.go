package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeSenna/golang-api-test/config"
	"github.com/guilhermeSenna/golang-api-test/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.CustomerRoute(router)
	router.Run(":8080")
}