package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/mt-inside/go-usvc"
	greeterapiv1 "github.com/mt-inside/greeter-sdk-go/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/mt-inside/mygreeter/pkg/health"
	"github.com/mt-inside/mygreeter/version"
)

var globalGreeting = "hello" // FIXME

func main() {
	log := usvc.GetLogger(false, 10)
	log.Info("Starting", "app", "mygreeter", "version", version.Version)

	go health.ListenHealthChecks(":8090")

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error(err, "Failed to listen")
		os.Exit(1)
	}

	server := grpc.NewServer()
	greeterapiv1.RegisterGreeterServiceServer(server, &greeterService{})
	greeterapiv1.RegisterConfigServiceServer(server, &configService{})
	reflection.Register(server)
	log.Info("Listening", "addr", l.Addr())
	err = server.Serve(l)
	if err != nil {
		log.Error(err, "Failed to serve gRPC")
		os.Exit(1)
	}
}

type greeterService struct {
	greeterapiv1.UnimplementedGreeterServiceServer
}

func (g *greeterService) Greet(ctx context.Context, req *greeterapiv1.GreetRequest) (*greeterapiv1.GreetResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &greeterapiv1.GreetResponse{Message: fmt.Sprintf("%s %s", globalGreeting, req.GetName())}, nil

}

type configService struct {
	greeterapiv1.UnimplementedConfigServiceServer
}

func (c *configService) GetGreeting(ctx context.Context, req *greeterapiv1.GetGreetingRequest) (*greeterapiv1.GetGreetingResponse, error) {
	return &greeterapiv1.GetGreetingResponse{Greeting: globalGreeting}, nil
}
func (c *configService) SetGreeting(ctx context.Context, req *greeterapiv1.SetGreetingRequest) (*greeterapiv1.SetGreetingResponse, error) {
	globalGreeting = req.GetGreeting()
	return &greeterapiv1.SetGreetingResponse{}, nil
}
