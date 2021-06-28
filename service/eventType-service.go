package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Crypto-Brothers/poc-vehicle-event-ledger-api/model"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type EventTypeService interface {
	Save(model.EventType) model.EventType
	FindAll() []model.EventType
}

type eventTypeService struct {
	eventTypes []model.EventType
}

func NewType() EventTypeService {
	return &eventTypeService{
		eventTypes: []model.EventType{},
	}
}

func (service *eventTypeService) Save(eventType model.EventType) model.EventType {

	var client = GetHederaClient()

	myTopicId, err := hedera.TopicIDFromString(os.Getenv("EVENT_TYPE_TOPIC_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("The topic ID = %v\n", myTopicId)

	ma, err := json.Marshal(eventType)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ma))

	//Create the transaction
	transaction := hedera.NewTopicMessageSubmitTransaction().
		SetTopicID(myTopicId).
		SetMessage([]byte(string(ma)))

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

	//Get the transaction consensus status
	transactionStatus := transactionReceipt.Status

	fmt.Printf("The transaction consensus status is %v\n", transactionStatus)
	//v2.0.0

	return eventType
}

func (service *eventTypeService) FindAll() []model.EventType {
	var client = GetHederaClient()

	myTopicId, err := hedera.TopicIDFromString(os.Getenv("EVENT_TYPE_TOPIC_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("The topic ID = %v\n", myTopicId)

	var results []model.EventType

	sub, err := hedera.NewTopicMessageQuery().
		SetTopicID(myTopicId).
		SetStartTime(time.Unix(0, 0)).
		Subscribe(client, func(message hedera.TopicMessage) {
			var ma model.EventType
			err := json.Unmarshal(message.Contents, &ma)
			if err != nil {
				println(err.Error(), ": error Unmarshalling")
			}
			fmt.Println(ma.Category, "-", ma.Type)
			results = append(results, ma)
		})

	if err != nil {
		println(err.Error(), ": error subscribing to the topic")
		return results
	}

	time.Sleep(3 * time.Second)
	sub.Unsubscribe()

	if err != nil {
		panic(err)
	}

	return results
}
