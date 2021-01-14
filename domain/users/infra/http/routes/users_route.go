package routes

import (
	userController "github.com/martinsd3v/simple-golang-project/domain/users/infra/http/controllers"

	"github.com/gin-gonic/gin"
)

//Users responsible for manage user routes
func (router *Router) Users(engine *gin.Engine) {

	controller := userController.Controller{Connection: router.Connection}

	group := engine.Group("/users")
	{
		group.POST("/login", controller.Auth)
		group.GET("/", controller.Show)
		group.GET("/:uuid", controller.Index)
		group.POST("/", controller.Create)
		group.PUT("/:uuid", controller.Update)
		group.DELETE("/:uuid", controller.Destroy)

	}
}
