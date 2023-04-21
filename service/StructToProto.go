package service

import (
	"FleetService/models"

	pb "github.com/Y-Fleet/Grpc-Api/api"

	"github.com/golang/protobuf/ptypes"
)

func StructToProto(vehicles []*models.Vehicle) *pb.VehicleListResponse {
	vehicleList := &pb.VehicleListResponse{}

	for _, vehicle := range vehicles {
		fuelConsumptionList := []*pb.FuelConsumption{}
		for _, fc := range vehicle.FuelConsumption {
			timestamp, _ := ptypes.TimestampProto(fc.Date)
			pbFuelConsumption := &pb.FuelConsumption{
				Date:             timestamp,
				DistanceTraveled: int32(fc.DistanceTraveled),
				FuelUsed:         int32(fc.FuelUsed),
				Consumption:      int32(fc.Consumption),
			}
			fuelConsumptionList = append(fuelConsumptionList, pbFuelConsumption)
		}

		maintenanceHistoryList := []*pb.MaintenanceHistory{}
		for _, mh := range vehicle.MaintenanceHistory {
			timestamp, _ := ptypes.TimestampProto(mh.Date)
			pbMaintenanceHistory := &pb.MaintenanceHistory{
				Date:        timestamp,
				Description: mh.Description,
				Cost:        int32(mh.Cost),
			}
			maintenanceHistoryList = append(maintenanceHistoryList, pbMaintenanceHistory)
		}

		pbAssignedDriver := &pb.AssignedDriver{
			Id:             vehicle.AssignedDriver.ID,
			Certifications: vehicle.AssignedDriver.Certifications,
		}

		pbVehicle := &pb.Vehicle{
			Id:                 vehicle.ID,
			Make:               vehicle.Make,
			Model:              vehicle.Model,
			Year:               int32(vehicle.Year),
			LicensePlateNumber: vehicle.LicensePlateNumber,
			VehicleIdNumber:    vehicle.VehicleIDNumber,
			WarehouseId:        vehicle.WarehouseID,
			Status:             vehicle.Status,
			FuelConsumption:    fuelConsumptionList,
			MaintenanceHistory: maintenanceHistoryList,
			AssignedDriver:     pbAssignedDriver,
		}

		vehicleList.Vehicles = append(vehicleList.Vehicles, pbVehicle)
	}

	return vehicleList
}
