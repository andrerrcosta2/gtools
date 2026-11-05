// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import "github.com/andrerrcosta2/gtools/core/seeders/random"

// DeepPlanet Representing the solar system hierarchy.
type DeepPlanet struct {
	Name       string  // Name of the planet
	DiameterKM float64 // Diameter in kilometers
	HasLife    bool    // Indicates if the planet supports life
}

func DeepPlanetAsValue(name string, diameterKM float64, hasLife bool) DeepPlanet {
	return DeepPlanet{
		Name:       name,
		DiameterKM: diameterKM,
		HasLife:    hasLife,
	}
}

func DeepPlanetAsRef(name string, diameterKM float64, hasLife bool) *DeepPlanet {
	return &DeepPlanet{
		Name:       name,
		DiameterKM: diameterKM,
		HasLife:    hasLife,
	}
}

func DeepPlanetAsRandValue() DeepPlanet {
	return DeepPlanet{
		Name:       random.Alphanumeric(1, 2, 30).At(0),
		DiameterKM: random.SingleOf[float64](),
		HasLife:    random.SingleOf[bool](),
	}
}

func DeepPlanetAsRandRef() *DeepPlanet {
	return &DeepPlanet{
		Name:       random.Alphanumeric(1, 2, 30).At(0),
		DiameterKM: random.SingleOf[float64](),
		HasLife:    random.SingleOf[bool](),
	}
}

type DeepStar struct {
	Name    string       // Name of the star
	MassKG  float64      // Mass in kilograms
	Planets []DeepPlanet // List of planets orbiting the star
}

func DeepStarAsValue(name string, massKG float64, planets []DeepPlanet) DeepStar {
	return DeepStar{
		Name:    name,
		MassKG:  massKG,
		Planets: planets,
	}
}

func DeepStarAsRef(name string, massKG float64, planets []DeepPlanet) *DeepStar {
	return &DeepStar{
		Name:    name,
		MassKG:  massKG,
		Planets: planets,
	}
}

func DeepStarAsRandValue() DeepStar {
	star := DeepStar{
		Name:   random.Alphanumeric(1, 2, 30).At(0),
		MassKG: random.SingleOf[float64](),
	}
	size := random.Int(1, 1, 10).At(0)
	planets := make([]DeepPlanet, size)
	for i := 0; i < size; i++ {
		planets[i] = DeepPlanetAsRandValue()
	}
	star.Planets = planets
	return star
}

func DeepStarAsRandRef() *DeepStar {
	star := &DeepStar{
		Name:   random.Alphanumeric(1, 2, 30).At(0),
		MassKG: random.SingleOf[float64](),
	}
	size := random.Int(1, 1, 10).At(0)
	planets := make([]DeepPlanet, size)
	for i := 0; i < size; i++ {
		planets[i] = DeepPlanetAsRandValue()
	}
	star.Planets = planets
	return star
}

type DeepGalaxy struct {
	Name      string   // Name of the galaxy
	StarCount int      // Total stars in the galaxy
	MainStar  DeepStar // A representative main star
	BlackHole string   // Name of the central black hole
}

func DeepGalaxyAsValue(name string, starCount int, mainStar DeepStar, blackHole string) DeepGalaxy {
	return DeepGalaxy{
		Name:      name,
		StarCount: starCount,
		MainStar:  mainStar,
		BlackHole: blackHole,
	}
}

func DeepGalaxyAsRef(name string, starCount int, mainStar DeepStar, blackHole string) *DeepGalaxy {
	return &DeepGalaxy{
		Name:      name,
		StarCount: starCount,
		MainStar:  mainStar,
		BlackHole: blackHole,
	}
}

func DeepGalaxyAsRandValue() DeepGalaxy {
	galaxy := DeepGalaxy{
		Name:      random.Alphanumeric(1, 2, 30).At(0),
		StarCount: random.Int(1, 1, 1000000).At(0),
		MainStar:  DeepStarAsRandValue(),
		BlackHole: random.Alphanumeric(1, 2, 30).At(0),
	}
	return galaxy
}

func DeepGalaxyAsRandRef() *DeepGalaxy {
	galaxy := &DeepGalaxy{
		Name:      random.Alphanumeric(1, 2, 30).At(0),
		StarCount: random.Int(1, 1, 1000000).At(0),
		MainStar:  DeepStarAsRandValue(),
		BlackHole: random.Alphanumeric(1, 2, 30).At(0),
	}
	return galaxy
}

type DeepUniverse struct {
	Name     string       // Universe name or ID
	AgeYears int64        // Age of the universe in years
	Galaxies []DeepGalaxy // List of galaxies in this universe
}

func DeepUniverseAsValue(name string, ageYears int64, galaxies []DeepGalaxy) DeepUniverse {
	return DeepUniverse{
		Name:     name,
		AgeYears: ageYears,
		Galaxies: galaxies,
	}
}

func DeepUniverseAsRef(name string, ageYears int64, galaxies []DeepGalaxy) *DeepUniverse {
	return &DeepUniverse{
		Name:     name,
		AgeYears: ageYears,
		Galaxies: galaxies,
	}
}

func DeepUniverseAsRandValue() DeepUniverse {
	universe := DeepUniverse{
		Name:     random.Alphanumeric(1, 2, 30).At(0),
		AgeYears: random.Int64(1, 1, 10000000000).At(0),
	}
	size := random.Int(1, 1, 10).At(0)
	galaxies := make([]DeepGalaxy, size)
	for i := 0; i < size; i++ {
		galaxies[i] = DeepGalaxyAsRandValue()
	}
	universe.Galaxies = galaxies
	return universe
}

func DeepUniverseAsRandRef() *DeepUniverse {
	universe := &DeepUniverse{
		Name:     random.Alphanumeric(1, 2, 30).At(0),
		AgeYears: random.Int64(1, 1, 10000000000).At(0),
	}
	size := random.Int(1, 1, 10).At(0)
	galaxies := make([]DeepGalaxy, size)
	for i := 0; i < size; i++ {
		galaxies[i] = DeepGalaxyAsRandValue()
	}
	universe.Galaxies = galaxies
	return universe
}
