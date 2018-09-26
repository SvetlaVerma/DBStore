package parser

import (
	"bufio"
	pb "dbstore/MicroServices/grpc"
	"encoding/csv"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

type Person struct {
	ID string   `json:"firstname"`
	Name  string   `json:"lastname"`
	Email   string `json:"email"`
	MobileNumber string `json:"mobilenumber"`
}

func openCSV(path string) (csvFile *os.File, err error) {

	csvFile, err = os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening csv file: %s", err)
	}
	return
}

//TODO move deleteRecord into a separate package
func DeleteRecord(recordID string, conn * grpc.ClientConn) (err error) {
	//and now delete
	//TODO: move the delete to another package
	dbClient := pb.NewRecordsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	deleteRequest := &pb.DeleteRequest{ID: recordID}
	defer cancel()
	_, err = dbClient.Delete(ctx, deleteRequest)
	if err != nil {
		return fmt.Errorf("error in deleting ID %s: %s", deleteRequest.ID, err)
	}
	log.Printf("ID %s successfully deleted", deleteRequest.ID)
	return
}

func Parse(path string, conn * grpc.ClientConn) (err error) {

	csvFile, err := openCSV(path)
	if err != nil {
		return fmt.Errorf("error parsing file: %s", err)
	}
	fmt.Println("opened file")
	dbClient := pb.NewRecordsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			return fmt.Errorf("error encountered when reading the file: %s", error)
		}
		record := &pb.Record{
			ID: line[0],
			Name: line[1],
			Email: line[2],
			MobileNumber: line[3],
		}
		storeRequest := &pb.StoreRequest{Record: record}
		defer cancel()
		_, err = dbClient.Store(ctx, storeRequest)
		if err != nil {
			return fmt.Errorf("error in storing ID %s: %s", storeRequest.Record.ID, err)
		}
		log.Printf("ID %s successfully stored", storeRequest.Record.ID)
	}
	return
}
