/*
 * the DBService is a micro service that stores/updates records to mongoDB
 */

package main

import (
	"dbstore/MicroServices/CSVParser/parser"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"path/filepath"
)

const (
	address     = "localhost:50053"
	CSVFilePath = "$PWD"
)

func main() {
	// Set up a connection to the DBService.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	path := filepath.Join(os.ExpandEnv(CSVFilePath), "parser/data.csv")

	//now parse the file.
	err = parser.Parse(path, conn)
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	//TODO move deletion to another package.
	fmt.Printf("========================Time to delete one of the record===============================")
	err = parser.DeleteRecord("100", conn)
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}
}
