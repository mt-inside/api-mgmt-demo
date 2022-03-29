package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mt-inside/go-usvc"
	greeterapiv1 "github.com/mt-inside/greeter-sdk-go/api/v1"
	"google.golang.org/grpc"

	"github.com/mt-inside/greetclient/version"
)

func main() {
	log := usvc.GetLogger(false, 10)
	log.Info("Starting", "app", "greetclient", "version", version.Version)

	conn, err := grpc.Dial("localhost:8080", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Error(err, "Failed to connect to server", "addr", "localhost:8080")
		os.Exit(1)
	}
	log.Info("Connected", "addr", conn.Target())

	greeter := greeterapiv1.NewGreeterServiceClient(conn)
	resp, err := greeter.Greet(
		context.Background(),
		&greeterapiv1.GreetRequest{
			Name: "matt",
		},
	)
	if err != nil {
		log.Error(err, "Failed to get greeted")
		os.Exit(1)
	}

	fmt.Println(resp.Message)
}
