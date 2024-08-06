package server

import (
	"log"
	"net"

	"github.com/ratheeshkumar/restaurant_user_serviceV1/internal/handlers"
	pb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(handlr *handlers.UserHandler) {
	log.Println("Connecting gRPC server")
	lis, err := net.Listen("tcp", ":8082")

	if err != nil {
		log.Fatal("error creating listener on port 8082")
	}

	grp := grpc.NewServer()
	pb.RegisterUserServicesServer(grp, handlr.UserServicesServer)

	log.Printf("listening on gRPC server 8082")
	if err := grp.Serve(lis); err != nil {
		log.Fatal("error connecting to gRPC server")

	}
}
