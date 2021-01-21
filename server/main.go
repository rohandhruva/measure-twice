//go:generate protoc -I src/main/proto --go_out=plugins=grpc:. echo_service.proto
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	echo "github.com/alimate/measurement/g/grpc"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) Echo(ctx context.Context, msg *echo.Message) (*echo.Message, error) {
	return msg, nil
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	e := Server{}

	grpcServer := grpc.NewServer()
	echo.RegisterEchoServiceServer(grpcServer, &e)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		fmt.Println("I'm listening to port 9000")
	}
}

