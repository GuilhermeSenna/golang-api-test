package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/guilhermeSenna/golang-api-test/controller"
)


func CustomerRoute(router *gin.Engine) {
	// Returns an array of up to 50 customers
	router.GET("/customers", controller.GetCustomers)

	// Returns a specific customer
	router.GET("/customers/:id", controller.GetCustomer)

	// Creates a new customer
	router.POST("/customers", controller.CreateCustomer)

	// Updates a customer
	router.PUT("/customers/:id", controller.UpdateCustomer)
}