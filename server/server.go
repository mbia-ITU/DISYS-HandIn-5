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

}

type Server struct {
	service.UnimplementedServiceServer
}

func OpenServer() {
	log.Print("Loading...")

	listener, err := net.Listen("tcp", "localhost:5000")

	if err != nil {
		log.Fatalf("Could not listen @ %s", err)
		return
	}

	log.Print("Server is setup at port 5000.")

	s := service.Server{}

	grpcServer := grpc.NewServer()

	service.RegisterServicesServer(grpcServer, &s)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Failed to start gRPC Server :: %v", err)
	}
}
