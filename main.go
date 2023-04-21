package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	service "FleetService/service"

	pb "github.com/Yfleet/shared_proto/api"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedFleetServiceServer
}

func (s *server) GetFleet(ctx context.Context, req *pb.Empty) (*pb.VehicleListResponse, error) {
	client, cancel := ConToDb()
	response, err := service.GetVehicle(client)
	cancel()
	return response, err
}

func ConToDb() (*mongo.Client, context.CancelFunc) {
	const connectionString = "mongodb://Fleet:Fleet@localhost:27022"
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	fmt.Println("Connected successfully to MongoDB")
	return client, cancel
}

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFleetServiceServer(s, &server{})
	reflection.Register(s)
	log.Println("Starting microservice on :50054")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
