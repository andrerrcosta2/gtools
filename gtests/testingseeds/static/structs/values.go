// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package structs

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static"
	"math/rand"
	"slices"
)

func allValues() []any {
	return slices.Concat(ofNameds(), ofComparisons(), ofModels(), ofImplementations(), ofDeeps())
}

func ofNameds() []any {
	return []any{
		BoolAsValue(), IntAsValue(), UintAsValue(), FloatAsValue(), ComplexAsValue(), StringAsValue(),
		SliceAsValue[any](), ChannelAsValue(), MapAsValue[any, any](),
	}
}

func ofComparisons() []any {
	return []any{
		NaturallyComparableAsValue(rand.Intn(1000000), "seed"),
		NaturallyComparableWithMethodsAsValue(rand.Intn(1000000), "seed"),
		NotComparableAsValue(rand.Intn(1000000), []string{"seed"}),
	}
}

func ofModels() []any {
	return []any{
		static.User{}, static.Account{}, static.Address{}, static.Credential{},
		static.Profile{},
	}
}

func ofImplementations() []any {
	return []any{
		CloserSuccessAsValue(), CloserErrorAsValue(), CloserReaderSuccessAsValue(),
		CloserReaderErrorAsValue(), CloserReaderWriterSuccessAsValue(), CloserReaderWriterErrorAsValue(),
		StringerAsValue("hello"), StringerBytesAsValue([]byte("hello")), StringerStringAsValue("hello"),
	}
}

func ofDeeps() []any {
	return []any{
		static.DeepUniverse{
			Name:     "",
			AgeYears: 1000000,
			Galaxies: []static.DeepGalaxy{
				{
					Name:      "Galaxy X",
					StarCount: 30000,
					MainStar: static.DeepStar{
						Name:   "Star X",
						MassKG: 0,
						Planets: []static.DeepPlanet{
							{Name: "Planet X", DiameterKM: 100000, HasLife: false},
							{Name: "Planet Y", DiameterKM: 500000, HasLife: false},
							{Name: "Planet Z", DiameterKM: 2600000, HasLife: false},
						},
					},
					BlackHole: "XXXXXX",
				},
				{
					Name:      "Galaxy Y",
					StarCount: 42000,
					MainStar: static.DeepStar{
						Name:   "Star Y",
						MassKG: 0,
						Planets: []static.DeepPlanet{
							{Name: "Planet A", DiameterKM: 100000, HasLife: false},
							{Name: "Planet B", DiameterKM: 500000, HasLife: false},
							{Name: "Planet C", DiameterKM: 2600000, HasLife: true},
						},
					},
					BlackHole: "YYYYYY",
				},
			},
		},
		static.DeepLibrary{
			Name:    "Library X",
			Address: "10th Street, New York",
			Sections: []static.DeepSection{
				{
					Name: "Science",
					Shelves: []static.DeepShelf{
						{
							Label: "A",
							Books: []static.DeepBook{
								{Title: "The Universe", Author: "John Doe", PageCount: 1000},
								{Title: "The Sun", Author: "John Doe", PageCount: 500},
								{Title: "The Earth", Author: "John Doe", PageCount: 200},
							},
						},
					},
				},
				{
					Name: "History",
					Shelves: []static.DeepShelf{
						{
							Label: "A",
							Books: []static.DeepBook{
								{Title: "The History of the World", Author: "John Doe", PageCount: 1000},
								{Title: "Ancient History", Author: "John Doe", PageCount: 500},
								{Title: "Wars", Author: "John Doe", PageCount: 200},
							},
						},
					},
				},
			},
		},
		static.DeepCompany{
			Name:         "Space X",
			Headquarters: "Space City",
			Warehouses: []static.DeepWarehouse{
				{
					Name:     "Warehouse A",
					Address:  "Space Street, Space City",
					Capacity: 1000000,
					ActiveTrucks: []static.DeepTruck{
						{
							LicensePlate: "CLP-123",
							DriverName:   "Alexander Smith",
							Packages: []static.DeepPackage{
								{TrackingID: "ABC-123", WeightKG: 10, IsFragile: false},
								{TrackingID: "DEF-456", WeightKG: 20, IsFragile: true},
								{TrackingID: "GHI-789", WeightKG: 30, IsFragile: true},
							},
						},
						{
							LicensePlate: "CLP-456",
							DriverName:   "John Doe",
							Packages: []static.DeepPackage{
								{TrackingID: "JKL-123", WeightKG: 40, IsFragile: false},
								{TrackingID: "MNO-456", WeightKG: 50, IsFragile: false},
								{TrackingID: "PQR-789", WeightKG: 60, IsFragile: true},
							},
						},
					},
				},
				{
					Name:     "Warehouse B",
					Address:  "Space Street, Space City",
					Capacity: 1000000,
					ActiveTrucks: []static.DeepTruck{
						{
							LicensePlate: "CLP-123",
							DriverName:   "Julia Carter",
							Packages: []static.DeepPackage{
								{TrackingID: "XYZ-123", WeightKG: 10, IsFragile: false},
								{TrackingID: "TUV-456", WeightKG: 20, IsFragile: true},
								{TrackingID: "QRS-789", WeightKG: 30, IsFragile: true},
							},
						},
						{
							LicensePlate: "CLP-456",
							DriverName:   "Martin Thomas",
							Packages: []static.DeepPackage{
								{TrackingID: "GHI-123", WeightKG: 40, IsFragile: false},
								{TrackingID: "JKL-456", WeightKG: 50, IsFragile: false},
								{TrackingID: "MNO-789", WeightKG: 60, IsFragile: true},
							},
						},
					},
				},
			},
		},
	}
}
