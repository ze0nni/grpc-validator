package main

import (
	"log"
	"net"

	"github.com/ze0nni/grpc-validator/api"

	"google.golang.org/grpc"

	"github.com/ze0nni/grpc-validator/server"
)

//go:generate protoc -I api\ api\proto\api.proto --go_out=plugins=grpc:api

const port = ":50051"

func main() {
	listener, err := net.Listen("tcp", port)
	if nil != err {
		log.Fatal(err)
	}
	grpsServer := grpc.NewServer()
	api.RegisterValidatorServer(grpsServer, server.NewValidatorServer())

	if err := grpsServer.Serve(listener); nil != err {
		log.Fatal(err)
	}
}
