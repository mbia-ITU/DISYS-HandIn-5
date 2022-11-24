package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	auctionTime = 60 * time.Second
)

func main() {

}

type Server struct {
	services.UnimplementedServiceServer
}

func OpenServer() {
	log.Print("Loading...")

	listener, err := net.Listen("tcp", "localhost:5000")

	if err != nil {
		log.Fatalf("Could not listen @ %s", err)
		return
	}

	log.Print("Server is setup at port 5000.")

	s := services.Server{}

	grpcServer := grpc.NewServer()

	services.RegisterServicesServer(grpcServer, &s)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Failed to start gRPC Server :: %v", err)
	}
}
