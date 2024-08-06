package menus

import (
	"log"

	menupb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/menus/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (menupb.MenuServiceClient, error) {
	grpc, err := grpc.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error dailing to gRPC cleint:8083")
		return nil, err
	}
	log.Printf("successfully connected to gRPC client :8083")
	return menupb.NewMenuServiceClient(grpc), nil
}
