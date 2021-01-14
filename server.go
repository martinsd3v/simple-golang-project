package main

import (
	"log"

	userRouter "github.com/martinsd3v/simple-golang-project/domain/users/infra/http/routes"
	infraData "github.com/martinsd3v/simple-golang-project/infra/data"
	httpMiddlewares "github.com/martinsd3v/simple-golang-project/infra/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {

	//Start a connection with database
	connection, err := infraData.SetupDB("./.env")

	if err != nil {
		panic(err)
	}

	//Register repositories on routes
	uRouter := userRouter.Router{Connection: connection}

	router := gin.Default()
	router.Use(httpMiddlewares.CORSMiddleware())
	uRouter.Users(router)

	router.Run(":8787")
}
