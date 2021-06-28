package main

import (
	"github.com/Crypto-Brothers/poc-vehicle-event-ledger-api/controller"
	"github.com/Crypto-Brothers/poc-vehicle-event-ledger-api/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	vehicleEventService    service.VehicleEventService       = service.NewEvent()
	vehicleEventController controller.VehicleEventController = controller.NewEvent(vehicleEventService)
	eventTypeService       service.EventTypeService          = service.NewType()
	eventTypeController    controller.EventTypeController    = controller.NewType(eventTypeService)
	vehicleTokenService    service.VehicleTokenService       = service.NewVehicleToken()
	vehicleTokenController controller.VehicleTokenController = controller.NewVehicleToken(vehicleTokenService)
)

func main() {
	router := gin.Default()
	// Enable CORS for requests UI domain (port)
	router.Use(cors.Default())

	router.GET("/vehicleEvents/:vin", func(ctx *gin.Context) {
		ctx.JSON(200, vehicleEventController.FindByVin(ctx))
	})

	router.GET("/vehicleEvents", func(ctx *gin.Context) {
		ctx.JSON(200, vehicleEventController.FindAll(ctx))
	})

	router.POST("/vehicleEvents", func(ctx *gin.Context) {
		ctx.JSON(200, vehicleEventController.Save(ctx))
	})

	router.GET("/eventTypes", func(ctx *gin.Context) {
		ctx.JSON(200, eventTypeController.FindAll())
	})

	router.POST("/eventTypes", func(ctx *gin.Context) {
		ctx.JSON(200, eventTypeController.Save(ctx))
	})

	router.POST("/vehicleToken", func(ctx *gin.Context) {
		println(ctx.Request.GetBody)
		ctx.JSON(200, vehicleTokenController.Tokenize(ctx))
	})

	router.Run(":8082")

}
