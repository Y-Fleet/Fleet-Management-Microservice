package models

import "time"

type FuelConsumption struct {
	Date             time.Time `bson:"date"`
	DistanceTraveled int       `bson:"distance_traveled"`
	FuelUsed         int       `bson:"fuel_used"`
	Consumption      int       `bson:"consumption"`
}

type MaintenanceHistory struct {
	Date        time.Time `bson:"date"`
	Description string    `bson:"description"`
	Cost        int       `bson:"cost"`
}

type AssignedDriver struct {
	ID             string   `bson:"_id"`
	Certifications []string `bson:"certifications"`
}

type Vehicle struct {
	ID                 string               `bson:"_id,omitempty"`
	Make               string               `bson:"make"`
	Model              string               `bson:"model"`
	Year               int                  `bson:"year"`
	LicensePlateNumber string               `bson:"license_plate_number"`
	VehicleIDNumber    string               `bson:"vehicle_id_number"`
	WarehouseID        string               `bson:"warehouse_id"`
	Status             string               `bson:"status"`
	FuelConsumption    []FuelConsumption    `bson:"fuel_consumption"`
	MaintenanceHistory []MaintenanceHistory `bson:"maintenance_history"`
	AssignedDriver     AssignedDriver       `bson:"assigned_driver"`
}
