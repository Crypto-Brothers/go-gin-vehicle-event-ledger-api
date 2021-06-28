# go-demo-ledger-rest

This sample app is a REST API that reads and writes messages to the Hedera distrubuted ledger.  It is written in Go and utilizes the [Hedera Go SDK](https://github.com/hashgraph/hedera-sdk-go). The function of this demo app is twofold, with the first function being to provide a programatic interface to create an asset NFT for vehicles, to establish clear ownership.  THe next fucntion is to log  maintenance/repairs events for those vehicles onto a publicly searchable, immutable ledger, in this case the Hedera network.  Owners could use it to prove that they performed regualar maintenance on their vehicle, which would be usefull for resale evalution purposes.  Orgazizations, like insurance companies and repair shops could, once authorized, also use it to record incidents like accident damage or regular maintenance activities.  Think of it as a decentralized [Carfax](https://www.carfax.com/vehicle-history-reports/).  The service providers that work on the vehicles are sourced from a Hedera topic that has first confirmed the validity of the provider.

## Setup

This sample app assumes you have already installed the GO distribution.  If not, you can find instructions [here](https://golang.org/doc/install)

Adiitionally, you will need at least 2 Hedera Portal profiles. To create your Hedera Portal profile register [here](https://portal.hedera.com/register).  Once registered, you'll need to note your Account ID and your Private Key.  These credential will be used by the the app to access any Hedera network services uned in the demo.

Before starting the project, create an .env file in the project root directory.  This file will store environemtn variable, such as your Hedera Account ID, your Private Key, and Topic IDs used by the app.

### Set Hedera Credentials

> .env
>
> NFT_ISSUER_ACCOUNT_ID= (set account id)
>
> NFT_ISSUER_PRIVATE_KEY= (set private key)
>
> VEHICLE_OWNER_ACCOUNT_ID= (set account id)
>
> VEHICLE_OWNER_PRIVATE_KEY= (set private key)
>
> VEHICLE_EVENT_TOPIC_ID= (set later)
>
> VERIFIED_SERVICER_TOPIC_ID= (set later)

This project writes messages to a Hedera pub/sub topic, so you will need to create a topic by executing the following command from the project root directory.

> go run setup/hederaTopicCreation.go

This will create a Hedera pub/sub topic and will return the Topic Ids for the application.
Edit the .env again and set the TOPIC_IDs

> .env
>
> NFT_ISSUER_ACCOUNT_ID= (set account id)
>
> NFT_ISSUER_PRIVATE_KEY= (set private key)
>
> VEHICLE_OWNER_ACCOUNT_ID= (set account id)
>
> VEHICLE_OWNER_PRIVATE_KEY= (set private key)
>
> VEHICLE_EVENT_TOPIC_ID= (set topic id)
>
> VERIFIED_SERVICER_TOPIC_ID= (set topic id))

Finally, execute the project.

> go run server.go

This will start a local webserver that serves the REST API used to create and read Hedera Topic messages.
THe default URL will be http://localhost:8082/vehicleEvents

Update server.go file to change the port, if desired.

## End Points

### Vehicle Token
These endpoints are related to the tokenization of a vehicle.

GET /vehicleToken - return all asset NFT info for a vehicle

POST /vehicleToken - create an asset NFT for a vehicle

PUT /vehicleToken - transfer ownership of the fixed asset NFT for a vehicle

### Vehicle Events
These end points are used to read and write to the Hedera Topic that will record all maintenace/repair events for each vehical.  For this demo, a topic for a particular vehicle will be created once an asset NFT for the vehical has been created.  The app will only allow the owner of the NFT and other authorized parties to record maintenace/reapir events to the topic.  This is to ensure the validity of the party recording the maintenace/reapir events.

GET /vehicleEvents - return all messages for the vehicle event topic

GET /vehicleEvents/[ :vin ] - return messages for the vehicle event topic filtered by VIN

POST /vehicleEvents/ - save vehicle event to topic

Expected JSON request format for POST
>
>      {
>
>        "vin": "GA94234351",
>  
>        "workdescription": "Oil Change & Tune Up",
>
>        "servicer": "Smith Auto Repair",
>
>        "technician": "Joe Smith",
>
>        "selectedfile": "receipt.jpg"
>
>      }


The expected "selectedfile" is an uploaded image for the receipt or work summary.  Eventually, this project will be expanded with functionality to upload this file to a distrubuted storage layer, like [IPFS](https://ipfs.io/). 
  
### Vehical Servicers
The next end points are to manage a topic containing verified servicers.  These would be verified businesses that service cars.  The actual verification is assumed to be done off ledger, wherein that process would subsequently call the Hedera API to create a topic message for the servicer, thus making it available to this app.

GET /verifiedServicer - return all messages for the verified servicer topic

POST /verifiedServicer/ - save vehicle event to topic

Expected JSON request format for POST
>
> {
>
>     "id": "1",
>
>     "name": "Jiffy Lube #241",
>
>     "streetaddress": "11 Oak Mill Rd.",
>
>     "city": "Mephis",
>
>     "postalcode": "54322",
>
>     "country": "USA",
>
>     "services": [
>
>         ["Oil Change", 20],
>         ["Brakes",120],
>         ["Trans Fluid Drain",75]
>
>     ],
>
>     "technicians": [
>
>         ["Alton Green", 1],
>         ["Raj Patel", 2],
>         ["Mary Cook", 3]
>
>     ]
>
> }
> 
## The UI  
The code for the UI used to interact with this REST API is in the [node-demo-ledger-ui repository](https://github.com/droatl2000/node-demo-ledger-ui)
