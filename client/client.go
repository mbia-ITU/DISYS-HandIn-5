package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	/*
		"sync"
		"time"
		"context"
		"fmt"
	*/

	service "github.com/mbia-ITU/DISYS-HandIn-5/gRPC/gRPC"
	/*
		"google.golang.org/grpc"
		"google.golang.org/protobuf/types/known/emptypb"
	*/)

var (
	bidderName  string
	rmDirectory = make(map[string]replicationManager)
)

type replicationManager struct {
	serviceClient service.ThisserviceClient
	address       string
}

func main() {

}

func getReplicationManagers() {

}

func auctionManager() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		replicationManagers = getReplicationManagers

		result, err := getHighestBid(&replicationManagers)
		if err != nil {
			log.Printf("Trouble fetching result from replication managers %d\n", err)
			continue
		}

		if scanner.Text() == "status" {
			if result.getStatus == service.Status_AUCTION_OVER {
				log.Printf("Auction is over. The winning bid was %v made by %v\n", result.Amount, result.Bidder)
				return
			} else {
				log.Printf("Current highest bid is %v made by %v.\n", result.Amount, result.Bidder)
				continue
			}

		}

		bid, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("Input %v could not be converted to integer: %v\n", bid, err)
		}

		if result.getStatus() == service.Status_AUCTION_OVER {
			log.Printf("Auction is over. The winning bid was %v made by %v\n", result.Amount, result.Bidder)
			return
		}

		if bid <= result.GetAmount {
			log.Printf("Tried to bid %v, but bid was to low compared to the current highest bid of %v made by %v\n", bid, result.Amount, result.Bidder)
			continue
		} else {
			log.Printf("Making a bid of %v. To beat the current highest bid.\n", bid)
			MakeABidToAllReplications(&replicationManagers, bid)
		}

	}
}

func MakeABid(bid int32) error {

	return nil

}

func MakeABidToAllReplications() {

}

func getHighestBid() {

}

func GetResult() {

}

func connectToNode() {

}

func connectToAllNodes() {

}
