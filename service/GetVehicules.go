package service

import (
	"FleetService/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	pb "github.com/Yfleet/shared_proto/api"
)

func GetVehicle(client *mongo.Client) (*pb.VehicleListResponse, error) {
	var vehicles []*models.Vehicle

	collection := client.Database("Fleet").Collection("Vehicules")
	filter := bson.M{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var Vehicle models.Vehicle
		err := cur.Decode(&Vehicle)
		if err != nil {
			return nil, err
		}
		vehicles = append(vehicles, &Vehicle)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	vehicleList := StructToProto(vehicles)

	return vehicleList, nil
}
