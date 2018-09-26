# Goal of this repo is to create microservices to parse a csv file and store the records read in MongoDB.

## Mongo start instructions
- Download mongodb macOS package using link https://docs.mongodb.com/manual/tutorial/install-mongodb-enterprise-on-os-x/

## Install MongoDB

## Start the Database
- open a terminal and type
  mongod

## Start the mongo shell (Needed if you want to quickly verify the record you have added programmatically)
- open a new terminal and type:
  mongo --host 127.0.0.1:27017

# Implementation
There are two micro services created:
- CSVParser to parse the given csv file. For each record it reads, it calls the grpc call to store the record that in turn calls
  the store handler function of DBService.

- DBService
  Stores the records into the MongoDB
  There are two packages created to implement the micro svc. The idea is even if the actual database under the hood changes
  (eg from MongoDB to map/postgresql, the front end of the service handler.go will not change)

## GRPC is the interface between the two micro services
- I have created add and update rpcs in db.proto but have implemented only Add rpc currently.

## How to run it:

### start the DBService (that interfaces with MongoDB)
- You can open two terminals. In one go to $HOME/go/src/DBStore/MicroServices/DBService and enter "go run main.go"

After successful start of service, run the CSVParser client micro service by doing:

### start the CSVParser by opening another terminal. go to $HOME/go/src/DBStore/MicroServices/CSVParser and enter "go run main.go"

The CSVParser does the following:
1) Parses the file data.csv and for each record, invoke the grpc "Store" request to the DBService
2) Deletes the last record (this is just for verifying delete functionality. needs to be moved to something else)