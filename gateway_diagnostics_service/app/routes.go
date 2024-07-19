package app

import (
	"golang_service_template/controllers"
	"golang_service_template/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	SERVICE_BASE_PATH      string
	internalRoutes         *gin.RouterGroup
	cognitoProtectedRoutes *gin.RouterGroup
)

func init() {

	SERVICE_BASE_PATH = os.Getenv("SERVICE_BASE_PATH")
	internalRoutes = Router.Group(SERVICE_BASE_PATH + "/internal")
	cognitoProtectedRoutes = Router.Group(SERVICE_BASE_PATH)

}

func SetupRoutesMiddleware() {

	internalRoutes.Use(middleware.LogRequest())
	internalRoutes.Use(middleware.AuthorizeApiKey())

	cognitoProtectedRoutes.Use(middleware.LogRequest())
	cognitoProtectedRoutes.Use(middleware.AuthorizeAwsToken())

}

func SetupHealthRoute() {

	Router.GET(SERVICE_BASE_PATH+"/v1/health", controllers.GetHealth)

}

func SetupInternalRoutes() {

	internalRoutes.GET("/users", controllers.GetUsers)
	internalRoutes.POST("/task/message", controllers.CreateMessage)

	internalRoutes.POST("/publish/ack", controllers.CreateAck)
	internalRoutes.POST("/publish/dcm", controllers.CreateDcm)

}

func SetupCognitoRoutes() {

	cognitoProtectedRoutes.GET("/users", controllers.GetUsers)

	cognitoProtectedRoutes.GET("/user/:user_id", controllers.GetUser)

	cognitoProtectedRoutes.DELETE("/user/:user_id", controllers.DeleteUser)

	cognitoProtectedRoutes.POST("/user", controllers.CreateUser)

	cognitoProtectedRoutes.GET("/devices", controllers.GetDevices)

	//this api shows how a meesage from a api can be written to a kafka topic
	cognitoProtectedRoutes.POST("/task", controllers.CreateTask)

	cognitoProtectedRoutes.POST("/insert/application", controllers.InsertApplication)

	cognitoProtectedRoutes.DELETE("/delete/application/{id}", controllers.DeleteApplication)

	cognitoProtectedRoutes.POST("/insert/info", controllers.InsertInfo)

	cognitoProtectedRoutes.POST("/delete/info/{id}", controllers.DeleteInfo)

	cognitoProtectedRoutes.POST("/messages", controllers.GetMessages)
}
