package app

import (
	"golang_service_template/clients"
	"golang_service_template/database"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var (
	Router *gin.Engine
)

func init() {

	gin.SetMode(gin.ReleaseMode)

	Router = gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("isHexadecimal", isHexadecimal)
		v.RegisterValidation("isSerialNumbers", isSerialNumbers)
	}
}

func StartApp() {

	defer func() {
		database.CloseDatabasePool()
	}()

	SetupHealthRoute()
	SetupRoutesMiddleware()
	SetupInternalRoutes()
	SetupCognitoRoutes()

	if err := database.InitializeDatabasePool(); err != nil {
		panic(err)
	}

	go func() {
		if err := Router.Run(":8083"); err != nil {
			panic(err)
		}
	}()

	//this kafka consumer will read from topic and call a api
	//go clients.KafkaConsumer("MyGroup1", "MyTopic1", services.KafkaConsumerCallback1)

	//this kafka consumer will read from topic and publish to 2 different topics using the same producer (the producer that reads from ToKafkaCh2 channel)
	//not yet done
	//go clients.KafkaConsumer("MyGroup2", "MyTopic2", services.KafkaConsumerCallback2)

	//in the routes.go file, /task api will write to kafka topic using this producer
	go clients.KafkaProducer(clients.ToKafkaCh1)

	//in the routes.go file, /task api will write to kafka topic using this producer
	//go clients.KafkaProducer(clients.ToKafkaCh2)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt

}
