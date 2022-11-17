package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeSenna/golang-api-test/config"
	"github.com/guilhermeSenna/golang-api-test/routes"
)
// import 

func main() {
	router := gin.New()
	config:Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}