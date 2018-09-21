package mongo

import (
	grpc "DBStore/MicroServices/grpc"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//Session encapsulates the params required for the new Mongo session
type Session struct {
	session* mgo.Session
	collection *mgo.Collection
}

type Record struct {
	ID       string    `bson:"id"`
	Name 	 string    `bson:"name"`
	Email     string    `bson:"email"`
	MobileNumber    string `bson:"mobile_snumber"`
}

//NewMongoSession encapsulates creation of a new mongo session
func NewMongoSession() (session* Session, err error) {
	session = &Session{}
	err = session.Init()
	if(err != nil) {
		fmt.Errorf("error in initializing mongoSession")
	}
	return
}

//Init initializes mongo session parameters
func (session* Session) Init() (err error) {
	Host := []string{
		"127.0.0.1:27017",
		// replica set addrs...
	}
	const (
		Database   = "myDB"
		Collection = "rand"
	)
	mgoSession, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
		// Username: Username,
		// Password: Password,
		// Database: Database,
		// DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
		// 	return tls.Dial("tcp", addr.String(), &tls.Config{})
		// },
	})
	if err != nil {
		return fmt.Errorf("error initializign mongoDB: %s", err)
	}
	session.session = mgoSession
	// Collection
	session.collection = mgoSession.DB(Database).C(Collection)
	fmt.Println("mongo session initialised. collection made")
	return
}

func (session* Session) addRecord(record *grpc.Record) (err error) {
	rec := Record{
		ID:       record.ID,
		Name: record.Name,
		Email:     record.Email,
		MobileNumber:    record.MobileNumber,
	}

	if err = session.collection.Insert(rec); err != nil {
		return fmt.Errorf("error in inserting record with name %s: %s", record.Name, err)
	}
	fmt.Printf("inserted record: %s\n", rec.ID)
	return
}

func (session* Session) updateRecord(record *grpc.Record) (err error) {
	fmt.Printf("record %s already exists", record.ID)

	var existingRecord Record
	err = session.collection.Find(bson.M{"id": record.ID}).One(&existingRecord)
	if err != nil {
		return fmt.Errorf("error in updating, record exists but error getting it: %s", err)
	}
	fmt.Printf("id %s, name %s, email %s mobilenumber %s", existingRecord.ID, existingRecord.Name, existingRecord.Email, existingRecord.MobileNumber)
	// Update
	selector := bson.M{"id": existingRecord.ID}
	updator := bson.M{"$set": bson.M{"id": record.ID, "name": record.Name, "email": record.Email, "mobile_snumbe": record.MobileNumber}}
	if err := session.collection.Update(selector, updator); err != nil {
		return fmt.Errorf("error in updating record %s; %s", record.ID, err)
	}

	fmt.Printf("id %s, name %s, email %s mobilenumber %s", existingRecord.ID, existingRecord.Name, existingRecord.Email, existingRecord.MobileNumber)
	return
}


func (session* Session) Add(record *grpc.Record) (err error) {
	fmt.Printf("mongo package Add. reached here: %s\n", record.ID)
	if (session.collection == nil) {
		log.Fatalf("collection fatal", err)
		return fmt.Errorf("collection nil" )
	}

	// check if the record already exists
	count, err := session.collection.Find(bson.M{"id": record.ID}).Count()
	if err != nil {
		return fmt.Errorf("error adding record %s while finding it: %s", record.ID,err)
	}
	fmt.Printf("number of records %d", count)
	//if not there, add it, else update it.
	if count == 0 {
		err = session.addRecord(record)
	} else {
		err = session.updateRecord(record)
	}
	var existingRecord Record
	//now check that it is really added/updated
	err = session.collection.Find(bson.M{"id": record.ID}).One(&existingRecord)
	if err != nil {
		return fmt.Errorf("error in Adding/updating record: %s", err)
	}
	return
}

func (session* Session) Update(id string, record *grpc.Record) (err error) {
	//TODO: implement update
	return
}