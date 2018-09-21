/*
 * the dbService is a micro service that stores/updates records to mongoDB
 */

package main

import (
	"dbstore/MicroServices/DBService/handler"
	pb "dbstore/MicroServices/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50052"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server, err := handler.NewServer()
	if err != nil {
		log.Fatalf("failed to create mongo session: %v", err)
	}
	pb.RegisterRecordsServer(s, server)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
