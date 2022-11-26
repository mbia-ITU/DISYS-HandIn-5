package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	service "github.com/mbia-ITU/DISYS-HandIn-5/gRPC/gRPC"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	auctionTime = 60 * time.Second
)

var (
	highestbid = service.Result{
		Bidder: "No one",
		Amount: 0,
		Status: service.Status_SUCCESS,
	}

	Locker sync.Mutex
)

func main() {
	go func() {
		time.Sleep(auctionTime)
		auctionOver()

		time.Sleep(3 * time.Second)
	}()

	OpenServer()
}

type Server struct {
	service.UnimplementedThisserviceServer
}

func OpenServer() {
	log.Print("Loading...")
	port := os.Getenv("PORT")
	address := fmt.Sprintf("localhost:%v", port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error while attempting to listen on port %v: %v", port, err)
		return
	}

	log.Print("Server is setup at port %v.", port)
	grpcServer := grpc.NewServer()

	server := Server{}
	service.RegisterThisserviceServer(grpcServer, &server)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}

func (s *Server) MakeABid(ctx context.Context, bid *service.Bid) (*service.Result, error) {
	Locker.Lock()
	defer Locker.Unlock()

	if highestbid.Status == service.Status_AUCTION_OVER {
		log.Printf("Tried to make a bid of %v by %v, but auction was closed.\n", bid.Amount, bid.Uid)
		return &service.Result{Status: service.Status_AUCTION_OVER}, nil
	}

	if bid.Amount <= highestbid.Amount {
		log.Printf("Tried to bid %v by %v, but bid was to low compared to the current highest bid of %v made by %v\n", bid.Amount, bid.Uid, highestbid.Amount, highestbid.Bidder)
		return &service.Result{Status: service.Status_TOO_LOW}, nil
	}

	highestbid = service.Result{
		Amount: bid.Amount,
		Bidder: bid.Uid,
		Status: service.Status_SUCCESS,
	}
	log.Printf("New highest bid of %v made by %v has beaten old bid of %v by %v\n", bid.Amount, bid.Uid, highestbid.Amount, highestbid.Bidder)

	return &highestbid, nil
}

func (s *Server) GetResult(ctx context.Context, _ *emptypb.Empty) (*service.Result, error) {
	Locker.Lock()
	defer Locker.Unlock()
	log.Printf("Highest bid is %v made by %v\n", highestbid.Amount, highestbid.Bidder)
	return &highestbid, nil
}

func auctionOver() {
	Locker.Lock()
	defer Locker.Unlock()

	highestbid.Status = service.Status_AUCTION_OVER
	log.Printf("Auction is over. The winning bid was %v made by %v\n", highestbid.Amount, highestbid.Bidder)
}
