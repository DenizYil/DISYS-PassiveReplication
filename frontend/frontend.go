package main

import (
	"context"
	"errors"
	"log"
	inc "main/increment"
	util "main/util"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type FrontEnd struct {
	primaryServer     inc.IncrementorClient
	primaryServerPort int
	inc.UnimplementedIncrementorServer
	lock sync.Mutex
}

func (fe *FrontEnd) Increment(ctx context.Context, request *inc.IncrementRequest) (*inc.IncrementResponse, error) {
	fe.lock.Lock()
	defer fe.lock.Unlock()

	context, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	log.Printf("Received request: %s", request)

	attempts := 0

	for {
		log.Println("Looping")
		if attempts == 3 {
			return nil, errors.New("PRIMARY_SERVER_UNRESPONSIVE")
		}

		attempts++

		if fe.primaryServer == nil {
			fe.SetPrimaryServer()

			if fe.primaryServer == nil {
				continue
			}
		}

		resp, err := fe.primaryServer.Increment(context, request)

		// Could be a new primary server was selected
		if err != nil {
			log.Printf("Can't seem to connect to the primary server.. Trying again. Attempts: %s", strconv.Itoa(attempts))
			fe.SetPrimaryServer()
		} else {
			return resp, nil
		}
	}
}

func (server *FrontEnd) Uptime(ctx context.Context, request *inc.UptimeRequest) (*inc.UptimeResponse, error) {
	return nil, errors.New("uptime is not supported through frontend")
}

func (fe *FrontEnd) SetPrimaryServer() {
	dat, _ := os.ReadFile("../port.txt")
	port, _ := strconv.Atoi(string(dat))

	if port == fe.primaryServerPort {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost:"+string(dat), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		fe.primaryServer = nil
		fe.primaryServerPort = -1
		return
	}

	fe.primaryServerPort = port
	fe.primaryServer = inc.NewIncrementorClient(conn)
}

func main() {
	log.Println("Loading FrontEnd now... Please wait!")

	port := strconv.Itoa(util.MinPort(1000))

	listener, err := net.Listen("tcp", "localhost:"+port)

	if err != nil {
		log.Fatalf("TCP failed to listen... %s", err)
	}

	log.Print("Listener registered - setting up server now...")

	grpcServer := grpc.NewServer()
	fe := &FrontEnd{}
	fe.SetPrimaryServer()

	inc.RegisterIncrementorServer(grpcServer, fe)

	log.Print("===============================================================================")
	log.Print(" ")
	log.Print("                      Welcome to the Incrementor Service!                      ")
	log.Print("                                    FrontEnd                                   ")
	log.Printf("                           Running on localhost:%s                            ", port)
	log.Print(" ")
	log.Print("===============================================================================")

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Failed to server gRPC serevr over port 9000")
	}
}
