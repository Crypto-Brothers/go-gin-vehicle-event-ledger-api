# go-demo-ledger-rest

This sample app is a REST API that reads and writes messages to the Hedera distrubuted ledger.  It is written in Go and utilizes the [Hedera Go SDK](https://github.com/hashgraph/hedera-sdk-go). The function of this demo app is twofold, with the first function being to provide a programatic interface to create an asset NFT for vehicles, to establish clear ownership.  THe next fucntion is to log  events such as ownership changers, usage summaries, maintenance/repairs events, damage, vehicle alerts for those vehicles onto a publicly searchable, immutable ledger, in this case the Hedera network.  All events would be logged by the vehicles IOT sensors or in the case of an ownership change, the authorized local tag offie.  Owners could use it to prove that they performed regualar maintenance on their vehicle, which would be usefull for resale evalution purposes.  Orgazizations, like insurance companies, also use it to retreive incidents like accident damage or driving habits.  THe manufacturer would have access to  performance and durability data.  Think of it as a decentralized and more exhaustive [Carfax](https://www.carfax.com/vehicle-history-reports/).  

## Setup

This sample app assumes you have already installed the GO distribution.  If not, you can find instructions [here](https://golang.org/doc/install)

Adiitionally, you will need at least 2 Hedera Portal profiles. To create your Hedera Portal profile register [here](https://portal.hedera.com/register).  Once registered, you'll need to note your Account ID and your Private Key.  These credential will be used by the the app to access any Hedera network services uned in the demo.

Before starting the project, create an .env file in the project root directory.  This file will store environemtn variable, such as your Hedera Account ID, your Private Key, and Topic IDs used by the app.

### Set Hedera Credentials

> .env
>
> ACCOUNT_ID= (set account id)
>
> PRIVATE_KEY= (set private key)
>
> EVENT_TYPE_TOPIC_ID= (set later)

This project writes messages to a Hedera pub/sub topic, so you will need to create a topic by executing the following command from the project root directory.

> go run setup/hederaTopicCreation.go

This will create a Hedera pub/sub topic and will return the Topic Ids for the application.
Edit the .env again and set the TOPIC_IDs

> .env
>
> ACCOUNT_ID= (set account id)
>
> PRIVATE_KEY= (set private key)
>
> EVENT_TYPE_TOPIC_ID= (set topic id))

Finally, execute the project.

> go run server.go

This will start a local webserver that serves the REST API used to create and read Hedera Topic messages.
THe default URL will be http://localhost:8082/vehicleEvents

Update server.go file to change the port, if desired.

## End Points

### Vehicle Token
These endpoints are related to the tokenization of a vehicle.

GET /vehicleToken - return all asset NFT info for a vehicle

POST /vehicleToken - create an asset NFT & a vehicle ledger for a vehicle

PUT /vehicleToken - transfer ownership of the fixed asset NFT for a vehicle

### Vehicle Events
These end points are used to read and write to the Hedera Topic that will record all life events for each vehicle.  For this demo, a topic for a particular vehicle will be created once an asset NFT for the vehical has been created.  The app will only allow the vehicle account and authorized parties, such as tax offices, to record events to the topic.

GET /vehicleEvents - return all messages for the vehicle event topic

GET /vehicleEvents/[ :vin ] - return messages for the vehicle event topic filtered by VIN

POST /vehicleEvents/ - save vehicle event to topic

Expected JSON request format for POST
```
    {
        "vin": "GA94234351",
        "eventcategory": "Vehicle Alerts",
        "eventtype": "Air Bags Deployed",
        "description": "Joe Smith",
        "relatefileName": "receipt.jpg"
    }
```

The expected "relatefileName" is an uploaded image for an optional receipt or image to supply context.  Eventually, this project will be expanded with functionality to upload this file to a distrubuted storage layer, like [IPFS](https://ipfs.io/). 
  
### Event Type
The next end points are to manage the event types.  

GET /verifiedServicer - return all messages for the verified servicer topic

POST /verifiedServicer/ - save vehicle event to topic

Expected JSON request format for POST
```
 {
    "eventcategory": "Ownership Change",
    "eventtypes": [
        ["Initial Purchase"],
        ["Transfer from Sale"],
        ["Repossesion"]
    ]
 }
```

## All event Types
### Vehicles are assumed to be EV / Smart Vehicles
```
Ownership Change
	Initial Purchase
	Transfer from Sale
	Repossesion
Compliance
	Emissions
Milage Milestone
	50K
	100K
	150K
    ...
Maintenance/Servicing
	Replace air filter
	Scheduled maintenance
	Electrical work
	New tires
	Battery replacement
	Brake work
	Fluid added/replaced
	Wheels aligned/balanced
	Other
Damage
	Accident
	Vandalism
	Weather
	Other
Usage Summary
	Self Driving Miles
	Manual Driving Miles
	Average Speed
	Max Speed
	Min Speed
	Speed Violations
	Lbs Towed
Vehicle Alerts
	Air Bags Deployed
	Check Engine Alert
	Battery Alert
	Brake Alert
	Other
```
## The UI  
The code for the UI used to interact with this REST API is in the [node-demo-ledger-ui repository](https://github.com/droatl2000/node-demo-ledger-ui)
