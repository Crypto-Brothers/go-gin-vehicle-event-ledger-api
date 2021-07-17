package controller

import (
	"github.com/Crypto-Brothers/go-gin-vehicle-event-ledger-api/model"
	"github.com/Crypto-Brothers/go-gin-vehicle-event-ledger-api/service"
	"github.com/gin-gonic/gin"
)

type VehicleTokenController interface {
	Tokenize(ctx *gin.Context) model.VehicleToken
}

type vehicleTokenController struct {
	service service.VehicleTokenService
}

func NewVehicleToken(service service.VehicleTokenService) VehicleTokenController {
	return &vehicleTokenController{
		service: service,
	}
}

func (c *vehicleTokenController) Tokenize(ctx *gin.Context) model.VehicleToken {
	var vehicleToken model.VehicleToken

	ctx.BindJSON(&vehicleToken)

	c.service.Tokenize(vehicleToken)
	return vehicleToken
}
