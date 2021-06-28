package controller

import (
	"github.com/Crypto-Brothers/poc-vehicle-event-ledger-api/model"
	"github.com/Crypto-Brothers/poc-vehicle-event-ledger-api/service"
	"github.com/gin-gonic/gin"
)

type EventTypeController interface {
	FindAll() []model.EventType
	Save(ctx *gin.Context) model.EventType
}

type eventTypeController struct {
	service service.EventTypeService
}

func NewType(service service.EventTypeService) EventTypeController {
	return &eventTypeController{
		service: service,
	}
}

func (c *eventTypeController) FindAll() []model.EventType {
	return c.service.FindAll()
}

func (c *eventTypeController) Save(ctx *gin.Context) model.EventType {
	var eventType model.EventType

	ctx.BindJSON(&eventType)

	c.service.Save(eventType)
	return eventType
}
