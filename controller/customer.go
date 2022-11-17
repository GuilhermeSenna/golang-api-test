 package controller 

 import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeSenna/golang-api-test/config"
	"github.com/guilhermeSenna/golang-api-test/models"
)

 func GetCustomer(c *gin.Context){
	var customer models.Customers
	config.DB.Where("id = ?", c.Param("id")).First(&customer)
	c.JSON(200, &customer)
 }

 func GetCustomers(c *gin.Context){
	customers := []models.Customers{}
	config.DB.Find(&customers)
	c.JSON(200, &customers)
 }

 func CreateCustomer(c *gin.Context){
	var customer models.Customers
	c.BindJSON(&customer)
	config.DB.Create(&customer)
	c.JSON(200, &customer)
 }

 func UpdateCustomer(c *gin.Context){
	var customer models.Customers
	config.DB.Where("id = ?", c.Param("id")).First(&customer)
	c.BindJSON(&customer)
	config.DB.Save(&customer)
	c.JSON(200, &customer)
 }