# Goal of this repo is to create microservices to parse a csv file and store it in the DB

# Implementation
To achieve the goal, there are two micro services created:
- CSVParser to parse the given csv file. The path of the file is injected to the file parser so it can then also be mocked
  for testing purposes.
    - TODO: need to pass dbclient as part of file parser initialization so then it can be mocked for testing purposes

- DBService
  Stores the records into the MongoDB (I missed the part where the test README said Database can be a map)
  There are two packages created to implement the micro svc. The idea is even if the actual database under the hood changes
  (eg from MongoDB to map/postgresql, the front end of the service handler.go will not change)

## GRPC is the interface between the two micro services
- I have created add and update rpcs in db.proto but have implemented only Add rpc currently.

## How to run it
-