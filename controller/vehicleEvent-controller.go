package controller

import (
	"github.com/Crypto-Brothers/go-gin-vehicle-event-ledger-api/model"
	"github.com/Crypto-Brothers/go-gin-vehicle-event-ledger-api/service"
	"github.com/gin-gonic/gin"
)

type VehicleEventController interface {
	FindByVin(ctx *gin.Context) []model.VehicleEvent
	FindAll(ctx *gin.Context) []model.VehicleEvent
	Save(ctx *gin.Context) model.VehicleEvent
}

type vehicleEventController struct {
	service service.VehicleEventService
}

func NewEvent(service service.VehicleEventService) VehicleEventController {
	return &vehicleEventController{
		service: service,
	}
}

func (c *vehicleEventController) FindAll(ctx *gin.Context) []model.VehicleEvent {
	return c.service.FindAll()
}

func (c *vehicleEventController) Save(ctx *gin.Context) model.VehicleEvent {
	var vehicleEvent model.VehicleEvent

	ctx.BindJSON(&vehicleEvent)

	c.service.Save(vehicleEvent)
	return vehicleEvent
}

func (c *vehicleEventController) FindByVin(ctx *gin.Context) []model.VehicleEvent {
	vin := ctx.Param("vin")

	return c.service.FindByVin(vin)
}
