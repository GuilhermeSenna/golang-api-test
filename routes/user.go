package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/guilhermeSenna/golang-api-test/controller"
)


func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}