package controllers

import (
	userDB "github.com/martinsd3v/simple-golang-project/domain/users/infra/data"
	userServices "github.com/martinsd3v/simple-golang-project/domain/users/services"
	toolsErrors "github.com/martinsd3v/simple-golang-project/tools/errors"
	toolsPaginator "github.com/martinsd3v/simple-golang-project/tools/paginator"

	"github.com/gin-gonic/gin"
	"github.com/martinsd3v/go-requestparser/parser"
)

//Show responsible for list all users
func (controller *Controller) Show(c *gin.Context) {
	paginatorDTO := toolsPaginator.PaginatorDTO{}
	parser.Parser(c.Request, &paginatorDTO)

	repo := userDB.SetupRepository(controller.Connection)
	results, paginator := userServices.ShowUserService(repo, paginatorDTO)

	if len(results.PublicUsers()) < 1 {
		response := toolsErrors.ErrorRequest{}
		response.Code = 400
		response.Message = toolsErrors.ErrorDestroyRegister
		c.JSON(response.Code, response)
	} else {
		response := toolsPaginator.ShowResultsDTO{}
		response.Results = results.PublicUsers()
		response.Paginator = paginator.Paginator
		c.JSON(200, response)
	}
}

//Index responsible for list unique users
func (controller *Controller) Index(c *gin.Context) {
	uuid := c.Param("uuid")

	repo := userDB.SetupRepository(controller.Connection)
	result := userServices.IndexUserService(repo, uuid)

	if result.UUID != "" {
		c.JSON(200, result.PublicUser())
	} else {
		response := toolsErrors.ErrorRequest{}
		response.Code = 400
		response.Message = toolsErrors.ErrorEmptyResults
		c.JSON(response.Code, response)
	}
}

//Create responsible for add a new user
func (controller *Controller) Create(c *gin.Context) {

	dto := userServices.CreateUserTDO{}
	parser.Parser(c.Request, &dto)

	repo := userDB.SetupRepository(controller.Connection)
	created, err := userServices.CreateUserService(repo, dto)

	if err.Message != "" {
		c.JSON(200, err)
	} else {
		c.JSON(400, created.PublicUser())
	}

	return
}

//Update responsible for edit a user
func (controller *Controller) Update(c *gin.Context) {
	uuid := c.Param("uuid")

	dto := userServices.UpdateUserTDO{}
	parser.Parser(c.Request, &dto)

	repo := userDB.SetupRepository(controller.Connection)
	updated, err := userServices.UpdateUserService(repo, dto, uuid)

	if err.Message != "" {
		c.JSON(200, err)
	} else {
		c.JSON(400, updated.PublicUser())
	}
}

//Destroy responsible for remove a user
func (controller *Controller) Destroy(c *gin.Context) {
	uuid := c.Param("uuid")
	response := toolsErrors.ErrorRequest{}

	repo := userDB.SetupRepository(controller.Connection)
	deleted := userServices.DestroyUserService(repo, uuid)

	if deleted {
		response.Code = 200
		response.Message = toolsErrors.SuccessDestroyRegister
	} else {
		response.Code = 400
		response.Message = toolsErrors.ErrorDestroyRegister
	}

	c.JSON(response.Code, response)
}
