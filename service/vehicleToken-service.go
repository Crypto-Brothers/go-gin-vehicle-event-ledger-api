package service

import (
	"fmt"

	"github.com/droatl2000/demo-ledger-rest/model"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type VehicleTokenService interface {
	Create(model.VehicleToken) model.VehicleToken
}

type vehicleTokenService struct {
	vehicleTokens []model.VehicleToken
}

func NewToken() VehicleTokenService {
	return &vehicleTokenService{
		vehicleTokens: []model.VehicleToken{},
	}
}

func (service *vehicleTokenService) Create(vehicleToken model.VehicleToken) model.VehicleToken {
	var client = GetHederaClient()

	//Create the transaction and freeze the unsigned transaction
	tokenCreateTransaction, err := hedera.NewTokenCreateTransaction().
		SetTokenName(vehicleToken.Name).
		SetTokenSymbol(vehicleToken.Symbol).
		SetTreasuryAccountID(vehicleToken.TreasuryId).
		SetInitialSupply(uint64(vehicleToken.InitialSupply)).
		SetAdminKey(vehicleToken.AdminKey).
		FreezeWith(client)

	if err != nil {
		panic(err)
	}

	//Sign with the admin private key of the token, sign with the token treasury private key, sign with the client operator private key and submit the transaction to a Hedera network
	txResponse, err := tokenCreateTransaction.Sign(vehicleToken.AdminKey).Sign(vehicleToken.TreasuryKey).Execute(client)

	if err != nil {
		panic(err)
	}

	//Request the receipt of the transaction
	receipt, err := txResponse.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	//Get the token ID from the receipt
	tokenId := *receipt.TokenID

	fmt.Printf("The new token ID is %v\n", tokenId)

	return vehicleToken
}
