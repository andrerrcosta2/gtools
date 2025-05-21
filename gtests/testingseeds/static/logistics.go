// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package static

// DeepPackage Representing a logistics system.
type DeepPackage struct {
	TrackingID string  // Unique ID for the package
	WeightKG   float64 // Weight in kilograms
	IsFragile  bool    // Whether the package is fragile
}

type DeepTruck struct {
	LicensePlate string        // Truck's license plate
	DriverName   string        // Name of the driver
	Packages     []DeepPackage // Packages loaded on the truck
}

type DeepWarehouse struct {
	Name         string      // Warehouse name
	Address      string      // Warehouse location
	Capacity     int         // Maximum number of trucks it can handle
	ActiveTrucks []DeepTruck // Trucks currently at the warehouse
}

type DeepCompany struct {
	Name         string          // Name of the company
	Headquarters string          // HQ address
	Warehouses   []DeepWarehouse // Warehouses owned by the company
}
