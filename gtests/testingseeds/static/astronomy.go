// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package static

// DeepPlanet Representing the solar system hierarchy.
type DeepPlanet struct {
	Name       string  // Name of the planet
	DiameterKM float64 // Diameter in kilometers
	HasLife    bool    // Indicates if the planet supports life
}

type DeepStar struct {
	Name    string       // Name of the star
	MassKG  float64      // Mass in kilograms
	Planets []DeepPlanet // List of planets orbiting the star
}

type DeepGalaxy struct {
	Name      string   // Name of the galaxy
	StarCount int      // Total stars in the galaxy
	MainStar  DeepStar // A representative main star
	BlackHole string   // Name of the central black hole
}

type DeepUniverse struct {
	Name     string       // Universe name or ID
	AgeYears int64        // Age of the universe in years
	Galaxies []DeepGalaxy // List of galaxies in this universe
}
