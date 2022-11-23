package main

import (
	"time"
)

const (
	auctionTime = 60 * time.Second
)

func main() {

}

type Server struct {
	services.UnimplementedServiceServer
}
