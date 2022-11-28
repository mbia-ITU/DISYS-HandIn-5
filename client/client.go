package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	service "github.com/mbia-ITU/DISYS-HandIn-5/gRPC/gRPC"
	"google.golang.org/grpc"
	//"google.golang.org/protobuf/types/known/emptypb"
)

var (
	bidderName  string
	rmDirectory = make(map[string]replicationManager)
	NumOfNodes  []string
)

type replicationManager struct {
	serviceClient service.ThisserviceClient
	address       string
}

func main() {

}

func getReplicationManagers() []replicationManager {

	replicationManagers := make([]replicationManager, 0)

	for _, nodeAddress := range NumOfNodes {
		if client, success := rmDirectory[nodeAddress]; success {
			replicationManagers = append(replicationManagers, client)
		} else {
			log.Printf("Reconnecting to: %v\n", nodeAddress)
			connect, err := getConnection(nodeAddress)

			if err == nil {
				client := service.NewThisserviceClient(connect)
				replicationManager := replicationManager{serviceClient: client, address: nodeAddress}
				replicationManagers = append(replicationManagers, replicationManager)
				rmDirectory[nodeAddress] = replicationManager
			} else {
				log.Printf("Attempted to reconnect to: %v. Could not succeed.\n", nodeAddress)
			}
		}
	}

	return replicationManagers
}

func auctionManager() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		replicationManagers := getReplicationManagers()

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

func connectToNode(addr string) (service.ThisserviceClient, error) {
	connection, err := getConnection(addr)
	if err != nil {
		log.Printf("Could not connect to node: %v", addr)
		return nil, err
	}

	client := service.NewThisserviceClient(connection)
	return client, nil
}

func connectToAllNodes() {
	servers := os.Getenv("SERVERS")
	numOfServers, err := strconv.Atoi(servers)
	if err != nil {
		log.Fatalf("Could not convert number of server: %s. To an integer with error: %d.", numOfServers, err)
	}

	if numOfServers < 1 {
		log.Fatalf("There are no servers. Exiting program.")
	}

	for i := 0; i < numOfServers; i++ {
		addr := fmt.Sprintf("biddingserver%d:5000", i+1)
		NumOfNodes = append(NumOfNodes, addr)
	}

	Locker := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, addr := range NumOfNodes {
		wg.Add(1)
		go func(addr string) {
			client, _ := connectToNode(addr)

			Locker.Lock()
			rmDirectory[addr] = replicationManager{serviceClient: client, address: addr}
			Locker.Unlock()
			wg.Done()
		}(addr)
	}
	wg.Wait()
}

func getConnection(addr string) (*grpc.ClientConn, error) {

	contextWithTimeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	connection, err := grpc.DialContext(contextWithTimeout, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to server: %s", err)
	}

	return connection, nil
}
