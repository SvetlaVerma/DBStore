package handler

import (
	"fmt"
	"golang.org/x/net/context"
	mongo "DBStore/MicroServices/DBService/mongo"
	rpc "DBStore/MicroServices/grpc"
	)

// Server is used to implement handler side code.
type Server struct{
	mgoSession *mongo.Session
}
//NewServer initializes Server
func NewServer() (server *Server, err error) {
	server = &Server{}
	server.mgoSession, err = mongo.NewMongoSession()
	if err != nil {
		return nil, fmt.Errorf("error in initializing mongoSession; %s", err)
	}
	fmt.Println("new server created")
	return
}

// Store adds a new record or updates an existing record
func (s *Server) Store(ctx context.Context, in *rpc.StoreRequest) (resp *rpc.StoreResponse, err error) {
	fmt.Printf("store request received for : %s\n", in.Record.ID)

	err = s.mgoSession.Add(in.Record)
	if err != nil {
		return nil, fmt.Errorf("error in adding mongo record id %s; %s", in.Record.ID, err)
	}
	return &rpc.StoreResponse{}, nil
}

// Delete deletes a record
func (s *Server) Delete(ctx context.Context, in *rpc.DeleteRequest) (resp *rpc.DeleteResponse, err error) {
	fmt.Printf("delete request received for : %s\n", in.ID)

	err = s.mgoSession.Delete(in.ID)
	if err != nil {
		return nil, fmt.Errorf("error in deleting mongo record id %s; %s", in.ID, err)
	}
	return &rpc.DeleteResponse{}, nil
}
