 package controller 

 import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeSenna/golang-api-test/config"
	"github.com/guilhermeSenna/golang-api-test/models"
	"github.com/guilhermeSenna/golang-api-test/utils"
	"encoding/json"
)


 func GetCustomer(c *gin.Context){
	var customer models.Customers
	err := config.DB.Where("id = ?", c.Param("id")).First(&customer).Error

	// Check if has generate an error with the query
	if err != nil {
		c.JSON(404, "Customer not found!")
		return
	}
	
	c.JSON(200, &customer) 
 }

 func GetCustomers(c *gin.Context){
	customers := []models.Customers{}
	// Limit to 50 users
	config.DB.Limit(50).Find(&customers)
	c.JSON(200, &customers)
 }

 func CreateCustomer(c *gin.Context){
	// Get customer parametrs
	var customer models.Customers
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&customer)

	// Hash password
	hashPassword, _ := utils.HashPassword(customer.Password)
	customerFormatted := models.Customers{
		First:     customer.First,
		Last:      customer.Last,
		Email:     customer.Email,
		Password:  hashPassword,
	}

	// Add to database
	config.DB.Create(&customerFormatted)
	c.JSON(201, customerFormatted)
 }

 func UpdateCustomer(c *gin.Context){

	// Get customer parametrs
	var customer models.Customers
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&customer)

	// Hash password
	hashPassword, _ := utils.HashPassword(customer.Password)
	customerFormatted := models.Customers{
		First:     customer.First,
		Last:      customer.Last,
		Email:     customer.Email,
		Password:  hashPassword,
	}

	// Add to database
	config.DB.Create(&customerFormatted)
	c.JSON(200, customerFormatted)
 }