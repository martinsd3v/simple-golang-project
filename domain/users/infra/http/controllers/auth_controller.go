package controllers

import (
	userDB "github.com/martinsd3v/simple-golang-project/domain/users/infra/data"
	userServices "github.com/martinsd3v/simple-golang-project/domain/users/services"

	"github.com/gin-gonic/gin"
	"github.com/martinsd3v/go-requestparser/parser"
)

//Auth responsible for auth user
func (controller *Controller) Auth(context *gin.Context) {
	dto := userServices.AuthenticateUserDTO{}
	parser.Parser(context.Request, &dto)

	repo := userDB.SetupRepository(controller.Connection)
	token, err := userServices.AuthenticateUserService(repo, dto)

	if err.Message != "" {
		context.JSON(200, err)
	} else {
		context.JSON(400, token)
	}

	return
}
