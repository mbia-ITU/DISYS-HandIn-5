package main

import (
	"log"
	"net"
	"time"

	service "github.com/mbia-ITU/DISYS-HandIn-5/gRPC/gRPC"
	"google.golang.org/grpc"
)

const (
	auctionTime = 60 * time.Second
)

func main() {

	OpenServer()
}

type Server struct {
	service.UnimplementedThisserviceServer
}

func OpenServer() {
	log.Print("Loading...")

	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalf("Error while attempting to listen on port 5000: %v", err)
		return
	}

	log.Print("Server is setup at port 5000.")
	grpcServer := grpc.NewServer()

	server := Server{}
	service.RegisterThisserviceServer(grpcServer, &server)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}
