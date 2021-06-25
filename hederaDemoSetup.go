package main

import (
	"fmt"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

func main() {
	//Loads the .env file and throws an error if it cannot load the variables from that file correctly
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("Unable to load environment variables from .env file. Error:\n%v\n", err))
	}

	//Grab your testnet account ID and private key from the .env file
	myAccountId, err := hedera.AccountIDFromString(os.Getenv("ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	myPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	//Print your testnet account ID and private key to the console to make sure there was no error
	fmt.Printf("The account ID is = %v\n", myAccountId)
	fmt.Printf("The private key is = %v\n", myPrivateKey)

	//Create your testnet client
	client := hedera.ClientForTestnet()
	client.SetOperator(myAccountId, myPrivateKey)

	createTopic(client, "Public Vehicle Event Ledger")
	createTopic(client, "Verified Vehicle Servicer Ledger")

	//Create the transaction and freeze the unsigned transaction
	tokenCreateTransaction, err := hedera.NewTokenCreateTransaction().
		SetTokenName("MaintLog Bucks").
		SetTokenSymbol("MLB").
		SetTreasuryAccountID(myAccountId).
		SetInitialSupply(10000).
		SetAdminKey(myPrivateKey).
		FreezeWith(client)
	if err != nil {
		panic(err)
	}

	// Sign with the admin private key of the token, sign with the token treasury private key,
	//  sign with the client operator private key and submit the transaction to a Hedera network
	txResponse, err := tokenCreateTransaction.Sign(myPrivateKey).Sign(myPrivateKey).Execute(client)
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
}

func createTopic(client *hedera.Client, topicName string) {
	//Create the transaction
	transaction := hedera.NewTopicCreateTransaction()
	transaction.SetTopicMemo(topicName)

	//Sign with the client operator private key and submit the transaction to a Hedera network
	txResponse, err := transaction.Execute(client)

	if err != nil {
		panic(err)
	}

	//Request the receipt of the transaction
	transactionReceipt, err := txResponse.GetReceipt(client)

	if err != nil {
		panic(err)
	}

	//Get the topic ID
	newTopicID := *transactionReceipt.TopicID

	fmt.Printf("%v topic ID is %v\n", topicName, newTopicID)
}
