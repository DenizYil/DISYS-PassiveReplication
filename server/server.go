package main

import (
	"context"
	"fmt"
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

type ServerSwag struct {
	lock   sync.Mutex
	value  int64
	uptime int64
	port   int64
	leader int64
	inc.UnimplementedIncrementorServer
	connections map[int]inc.IncrementorClient
}

var server *ServerSwag

func (server *ServerSwag) Uptime(ctx context.Context, request *inc.UptimeRequest) (*inc.UptimeResponse, error) {
	return &inc.UptimeResponse{Uptime: server.uptime}, nil
}

func (server *ServerSwag) Increment(ctx context.Context, request *inc.IncrementRequest) (*inc.IncrementResponse, error) {
	server.lock.Lock()
	defer server.lock.Unlock()

	log.Printf("Received request: %s", request)
	log.Printf("Server value: %s", strconv.Itoa(int(server.value)))

	if request.Value <= server.value {
		log.Printf("Request value was lower or equal to than the current value")
		return &inc.IncrementResponse{Success: false, Before: server.value}, nil
	}

	log.Printf("Request value has been changed to %s", strconv.Itoa(int(request.Value)))

	response := &inc.IncrementResponse{Success: true, Before: server.value}
	server.value = request.Value

	// I'm the leader, so I need to update everyone else
	if server.leader == server.port {
		log.Println("Updating all replicas now, as I am the leader")
		go func() {
			os.WriteFile("../value.txt", []byte(strconv.Itoa(int(server.value))), 0644)
		}()

		wg := sync.WaitGroup{}

		for k, v := range server.connections {

			wg.Add(1)

			go func(port int, client inc.IncrementorClient) {
				defer wg.Done()

				_, err := client.Increment(ctx, request)

				if err != nil {
					delete(server.connections, port)
				} else {
					log.Printf("Server with port %s has been updated", strconv.Itoa(port))
				}
			}(k, v)
		}

		wg.Wait()
	}

	return response, nil
}

func StartServer() {
	port := strconv.Itoa(util.MinPort(10000))

	log.Printf("Starting server now for port %s, please wait", port)

	listener, err := net.Listen("tcp", "localhost:"+port)

	if err != nil {
		log.Fatalf("TCP failed to listen... %s", err)
	}

	log.Print("Listener registered - setting up server now...")

	_port, _ := strconv.ParseInt(port, 10, 64)

	grpcServer := grpc.NewServer()
	server = &ServerSwag{
		uptime:      time.Now().UnixNano(),
		port:        _port,
		connections: make(map[int]inc.IncrementorClient),
	}

	inc.RegisterIncrementorServer(grpcServer, server)

	data, _ := os.ReadFile("value.txt")
	value, _ := strconv.ParseInt(string(data), 10, 64)
	server.value = value

	log.Print("===============================================================================")
	log.Print(" ")
	log.Print("                      Welcome to the Incrementor Service!                      ")
	log.Print("                                     Server                                    ")
	log.Printf("                           Running on localhost:%s                            ", port)
	log.Print(" ")
	log.Print("===============================================================================")

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Failed to server gRPC serevr over port 9000")
	}
}

func Connect(port int) inc.IncrementorClient {
	if client, ok := server.connections[port]; ok {
		return client
	}

	address := fmt.Sprintf("localhost:%v", port)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		return nil
	}

	client := inc.NewIncrementorClient(conn)

	server.connections[port] = client
	return client
}

func FindLeader() {
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	for {
		if server == nil {
			continue
		}

		uptime := server.uptime
		port := server.port

		for i := 1; i < 10; i++ {
			_port := i * 10000

			// We don't want our own server
			if _port == int(server.port) {
				continue
			}

			wg.Add(1)

			go func() {
				defer wg.Done()

				client := Connect(_port)

				if client == nil {
					return
				}

				resp, err := client.Uptime(context.Background(), &inc.UptimeRequest{})

				if err != nil {
					delete(server.connections, _port)
					return
				}

				lock.Lock()
				if resp.Uptime < uptime {
					port = int64(_port)
					uptime = resp.Uptime
				}
				lock.Unlock()
			}()
		}

		wg.Wait()

		// We were just made leader
		if server.port == port && server.leader != port {
			log.Printf("%s >> I am now the primary server!", strconv.Itoa(int(port)))
			go func() {
				os.WriteFile("../port.txt", []byte(strconv.Itoa(int(port))), 0644)
			}()
		}

		server.leader = port
	}
}

func main() {
	go StartServer()
	go FindLeader()

	select {}
}
