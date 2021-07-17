package service

import (
	"fmt"
	"os"

	"github.com/Crypto-Brothers/go-gin-vehicle-event-ledger-api/model"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type VehicleTokenService interface {
	Tokenize(model.VehicleToken) model.VehicleToken
}

type vehicleTokenService struct {
	vehicleTokens []model.VehicleToken
}

func NewVehicleToken() VehicleTokenService {
	return &vehicleTokenService{
		vehicleTokens: []model.VehicleToken{},
	}
}

func (service *vehicleTokenService) Tokenize(vehicleToken model.VehicleToken) model.VehicleToken {

	print("account:" + os.Getenv("ACCOUNT_ID"))

	//Grab your testnet account ID and private key from the .env file
	authAccountId, err := hedera.AccountIDFromString(os.Getenv("ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	authPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	var client = GetHederaClient()

	//Create the transaction and freeze the unsigned transaction
	tokenCreateTransaction, err := hedera.NewTokenCreateTransaction().
		SetTokenName(vehicleToken.Name).
		SetTokenSymbol(vehicleToken.Symbol).
		SetTreasuryAccountID(authAccountId).
		SetInitialSupply(1).
		SetAdminKey(authPrivateKey).
		FreezeWith(client)

	if err != nil {
		panic(err)
	}

	//Sign with the admin private key of the token, sign with the token treasury private key, sign with the client operator private key and submit the transaction to a Hedera network
	txResponseCreate, err := tokenCreateTransaction.Sign(vehicleToken.AdminKey).Sign(vehicleToken.TreasuryKey).Execute(client)

	if err != nil {
		panic(err)
	}

	//Request the receipt of the transaction
	receiptCreate, err := txResponseCreate.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	//Get the token ID from the receipt
	tokenId := *receiptCreate.TokenID

	fmt.Printf("The new token ID is %v\n", tokenId)

	var topicId = CreateTopic(client, tokenId.String()+"_EVENT_LOG")

	var tokenmemo = "{topicid:\"" + topicId.String() + "\"}"
	vehicleToken.Memo = tokenmemo

	tokenUpdateTransaction, err := hedera.NewTokenUpdateTransaction().
		SetTokenID(tokenId).
		SetTokenMemo(vehicleToken.Memo).
		FreezeWith(client)

	if err != nil {
		panic(err)
	}

	//Sign with the admin private key of the token, sign with the client operator private key and submit the transaction to a Hedera network
	txResponseUpdate, err := tokenUpdateTransaction.Sign(vehicleToken.AdminKey).Execute(client)

	if err != nil {
		panic(err)
	}

	//Request the receipt of the transaction
	receiptUpdate, err := txResponseUpdate.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	//Get the transaction consensus status
	statusUpdate := receiptUpdate.Status

	fmt.Printf("The update token transaction consensus status is %v\n", statusUpdate)

	return vehicleToken
}
