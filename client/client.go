package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	inc "main/increment"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client inc.IncrementorClient

func FindFrontEnd() {
	lock := sync.Mutex{}
	found := false
	foundCh := make(chan bool)

	for i := 0; i < 10; i++ {
		port := i * 1000
		go func() {
			address := fmt.Sprintf("localhost:%v", port)

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

			if err != nil || found {
				return
			}

			lock.Lock()
			client = inc.NewIncrementorClient(conn)
			found = true
			foundCh <- true
			lock.Unlock()
		}()
	}

	select {
	case <-foundCh:
	case <-time.After(2 * time.Second):
		return
	}
}

func Increment(value int64) string {
	attempts := 0

	for {
		if attempts == 2 {
			return "All front end servers seems to be down at the moment. \n" +
				"Make sure that you have at least one running in the background."
		}

		attempts++

		resp, err := client.Increment(context.Background(), &inc.IncrementRequest{Value: value})

		if err == nil {
			if resp.Success {
				return fmt.Sprintf("Cool! Incremented to %v - was %v before.", value, resp.Before)
			} else {
				return fmt.Sprintf("Uh oh! You cannot increment to %v, since the value is already %v.", value, resp.Before)
			}
		}

		if strings.Contains(err.Error(), "PRIMARY_SERVER_UNRESPONSIVE") {
			return "The primary server is somehow unresponsive. Please try again later"
		}

		FindFrontEnd()
	}
}

func main() {
	log.Printf("Loading client - please wait...")

	FindFrontEnd()

	if client == nil {
		log.Println("All front end servers seems to be down at the moment.")
		log.Println("Make sure that you have at least one running in the background.")
		return
	}

	log.Printf("Client loaded! Write your increment.")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		command := scanner.Text()
		input, err := strconv.ParseInt(command, 10, 64)
		if err != nil {
			log.Printf("%v is not an integer.\n", command)
		} else {
			log.Println(Increment(input))
		}
	}
}
