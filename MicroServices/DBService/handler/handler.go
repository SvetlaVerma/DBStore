package handler

import (
	"fmt"
	"golang.org/x/net/context"
	mongo "dbstore/MicroServices/DBService/mongo"
	rpc "dbstore/MicroServices/grpc"
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

// Store adds a new record
func (s *Server) Store(ctx context.Context, in *rpc.StoreRequest) (resp *rpc.StoreResponse, err error) {
	fmt.Printf("store request received for : %s\n", in.Record.ID)

	err = s.mgoSession.Add(in.Record)
	if err != nil {
		return nil, fmt.Errorf("error in adding mongo record id %s; %s", in.Record.ID, err)
	}
	return &rpc.StoreResponse{}, nil
}

//Update updates an existing store
func (s *Server) Update(ctx context.Context, in *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	return &rpc.UpdateResponse{}, nil
}