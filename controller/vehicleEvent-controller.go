package controller

import (
	"github.com/Crypto-Brothers/poc-vehicle-event-ledger-api/model"
	"github.com/Crypto-Brothers/poc-vehicle-event-ledger-api/service"
	"github.com/gin-gonic/gin"
	"github.com/hashgraph/hedera-sdk-go/v2"
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
	topicid, terr := hedera.TopicIDFromString(ctx.Param("topicid"))
	if terr != nil {
		panic(terr)
	}

	return c.service.FindAll(topicid)
}

func (c *vehicleEventController) Save(ctx *gin.Context) model.VehicleEvent {
	var vehicleEvent model.VehicleEvent

	ctx.BindJSON(&vehicleEvent)

	topicid, terr := hedera.TopicIDFromString(ctx.Param("topicid"))
	if terr != nil {
		panic(terr)
	}

	c.service.Save(topicid, vehicleEvent)
	return vehicleEvent
}

func (c *vehicleEventController) FindByVin(ctx *gin.Context) []model.VehicleEvent {

	topicid, terr := hedera.TopicIDFromString(ctx.Param("topicid"))
	if terr != nil {
		panic(terr)
	}
	vin := ctx.Param("vin")

	return c.service.FindByVin(topicid, vin)
}
