// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package structs

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static"
	"math/rand"
	"slices"
)

func allRefs() []any {
	return slices.Concat(ofNamedRefs(), ofComparisonRefs(), ofModelRefs(), ofImplementationRefs(), ofDeepsRefs())
}

func ofNamedRefs() []any {
	return []any{
		BoolAsRef(), IntAsRef(), UintAsRef(), FloatAsRef(), ComplexAsRef(), StringAsRef(),
		SliceAsRef[any](), ChannelAsRef(), MapAsRef[any, any](),
	}
}

func ofComparisonRefs() []any {
	return []any{
		NaturallyComparableAsRef(rand.Intn(1000000), "seed"),
		NaturallyComparableWithMethodsAsRef(rand.Intn(1000000), "seed"),
		NotComparableAsRef(rand.Intn(1000000), []string{"seed"}),
	}
}

func ofModelRefs() []any {
	return []any{
		static.User{}, static.Account{}, static.Address{}, static.Credential{},
		static.Profile{},
	}
}

func ofImplementationRefs() []any {
	return []any{
		CloserSuccessAsRef(), CloserErrorAsRef(), CloserReaderSuccessAsRef(),
		CloserReaderErrorAsRef(), CloserReaderWriterSuccessAsRef(), CloserReaderWriterErrorAsRef(),
		StringerAsRef("hello"), StringerBytesAsRef([]byte("hello")), StringerStringAsRef("hello"),
	}
}

func ofDeepsRefs() []any {
	return []any{
		&static.DeepUniverse{
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
		&static.DeepLibrary{
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
		&static.DeepCompany{
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
