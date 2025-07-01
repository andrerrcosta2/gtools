// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import "github.com/andrerrcosta2/gtools/core/seeders/random"

// DeepPackage Representing a logistics system.
type DeepPackage struct {
	TrackingID string  // Unique ID for the package
	WeightKG   float64 // Weight in kilograms
	IsFragile  bool    // Whether the package is fragile
}

func DeepPackageAsValue(trackingID string, weightKG float64, isFragile bool) DeepPackage {
	return DeepPackage{
		TrackingID: trackingID,
		WeightKG:   weightKG,
		IsFragile:  isFragile,
	}
}

func DeepPackageAsRef(trackingID string, weightKG float64, isFragile bool) *DeepPackage {
	return &DeepPackage{
		TrackingID: trackingID,
		WeightKG:   weightKG,
		IsFragile:  isFragile,
	}
}

func DeepPackageAsRandValue() DeepPackage {
	return DeepPackage{
		TrackingID: random.Alphanumeric(1, 2, 30).At(0),
		WeightKG:   random.SingleOf[float64](),
		IsFragile:  random.SingleOf[bool](),
	}
}

func DeepPackageAsRandRef() *DeepPackage {
	return &DeepPackage{
		TrackingID: random.Alphanumeric(1, 2, 30).At(0),
		WeightKG:   random.SingleOf[float64](),
		IsFragile:  random.SingleOf[bool](),
	}
}

type DeepTruck struct {
	LicensePlate string        // Truck's license plate
	DriverName   string        // Name of the driver
	Packages     []DeepPackage // Packages loaded on the truck
}

func DeepTruckAsValue(licensePlate, driverName string, packages []DeepPackage) DeepTruck {
	return DeepTruck{
		LicensePlate: licensePlate,
		DriverName:   driverName,
		Packages:     packages,
	}
}

func DeepTruckAsRef(licensePlate, driverName string, packages []DeepPackage) *DeepTruck {
	return &DeepTruck{
		LicensePlate: licensePlate,
		DriverName:   driverName,
		Packages:     packages,
	}
}

func DeepTruckAsRandValue() DeepTruck {
	truck := DeepTruck{
		LicensePlate: random.Alphanumeric(1, 2, 10).At(0),
		DriverName: random.Alphabet(1, 3, 15).At(0) + " " +
			random.Alphabet(1, 3, 15).At(0),
	}
	size := random.Int(1, 1, 5).At(0)
	packages := make([]DeepPackage, size)
	for i := 0; i < size; i++ {
		packages[i] = DeepPackageAsRandValue()
	}
	truck.Packages = packages
	return truck
}

func DeepTruckAsRandRef() *DeepTruck {
	truck := &DeepTruck{
		LicensePlate: random.Alphanumeric(1, 2, 10).At(0),
		DriverName: random.Alphabet(1, 3, 15).At(0) + " " +
			random.Alphabet(1, 3, 15).At(0),
	}
	size := random.Int(1, 1, 5).At(0)
	packages := make([]DeepPackage, size)
	for i := 0; i < size; i++ {
		packages[i] = DeepPackageAsRandValue()
	}
	truck.Packages = packages
	return truck
}

type DeepWarehouse struct {
	Name         string      // Warehouse name
	Address      string      // Warehouse location
	Capacity     int         // Maximum number of trucks it can handle
	ActiveTrucks []DeepTruck // Trucks currently at the warehouse
}

func DeepWarehouseAsValue(name, address string, capacity int, activeTrucks []DeepTruck) DeepWarehouse {
	return DeepWarehouse{
		Name:         name,
		Address:      address,
		Capacity:     capacity,
		ActiveTrucks: activeTrucks,
	}
}

func DeepWarehouseAsRef(name, address string, capacity int, activeTrucks []DeepTruck) *DeepWarehouse {
	return &DeepWarehouse{
		Name:         name,
		Address:      address,
		Capacity:     capacity,
		ActiveTrucks: activeTrucks,
	}
}

func DeepWarehouseAsRandValue() DeepWarehouse {
	warehouse := DeepWarehouse{
		Name:     random.Alphabet(1, 3, 15).At(0),
		Address:  random.Alphabet(1, 3, 30).At(0),
		Capacity: random.Int(1, 1, 100).At(0),
	}
	size := random.Int(1, 1, 20).At(0)
	activeTrucks := make([]DeepTruck, size)
	for i := 0; i < size; i++ {
		activeTrucks[i] = DeepTruckAsRandValue()
	}
	warehouse.ActiveTrucks = activeTrucks
	return warehouse
}

func DeepWarehouseAsRandRef() *DeepWarehouse {
	warehouse := &DeepWarehouse{
		Name:     random.Alphabet(1, 3, 15).At(0),
		Address:  random.Alphabet(1, 3, 30).At(0),
		Capacity: random.Int(1, 1, 100).At(0),
	}
	size := random.Int(1, 1, 20).At(0)
	activeTrucks := make([]DeepTruck, size)
	for i := 0; i < size; i++ {
		activeTrucks[i] = DeepTruckAsRandValue()
	}
	warehouse.ActiveTrucks = activeTrucks
	return warehouse
}

type DeepCompany struct {
	Name         string          // Name of the company
	Headquarters string          // HQ address
	Warehouses   []DeepWarehouse // Warehouses owned by the company
}

func DeepCompanyAsValue(name, headquarters string, warehouses []DeepWarehouse) DeepCompany {
	return DeepCompany{
		Name:         name,
		Headquarters: headquarters,
		Warehouses:   warehouses,
	}
}

func DeepCompanyAsRef(name, headquarters string, warehouses []DeepWarehouse) *DeepCompany {
	return &DeepCompany{
		Name:         name,
		Headquarters: headquarters,
		Warehouses:   warehouses,
	}
}

func DeepCompanyAsRandValue() DeepCompany {
	company := DeepCompany{
		Name:         random.Alphabet(1, 3, 15).At(0),
		Headquarters: random.Alphabet(1, 3, 30).At(0),
	}
	size := random.Int(1, 1, 10).At(0)
	warehouses := make([]DeepWarehouse, size)
	for i := 0; i < size; i++ {
		warehouses[i] = DeepWarehouseAsRandValue()
	}
	company.Warehouses = warehouses
	return company
}

func DeepCompanyAsRandRef() *DeepCompany {
	company := &DeepCompany{
		Name:         random.Alphabet(1, 3, 15).At(0),
		Headquarters: random.Alphabet(1, 3, 30).At(0),
	}
	size := random.Int(1, 1, 10).At(0)
	warehouses := make([]DeepWarehouse, size)
	for i := 0; i < size; i++ {
		warehouses[i] = DeepWarehouseAsRandValue()
	}
	company.Warehouses = warehouses
	return company
}
