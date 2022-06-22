package main

import (
	"log"
	"net"

	"github.com/ffigueiredo-sp/fullycle_go/services"
	"github.com/ffigueiredo-sp/fullycle_go_grpc/pb/pb"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserServvice())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
