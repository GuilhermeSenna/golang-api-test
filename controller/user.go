 package controller 

 import "github.com/gin-gonic/gin"

 func UserController(c *gin.Context){
	c.String(200, "Hello World!")
 }
