//Package space has Age function to calculate age of different planets
package space

// Planet is a string variable to store the planet name
type Planet string

// Age calculate the space age
func Age(seconds float64, planet Planet) float64 {
	earthRevolution := 31557600
	switch planet {
	// Mercury: orbital period 0.2408467 Earth years
	case "Mercury":
		return seconds / (0.2408467 * earthRevolution)
		// Venus: orbital period 0.61519726 Earth years
	case "Venus":
		return seconds / (0.61519726 * earthRevolution)
		// Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
	case "Earth":
		return seconds / (1.0 * earthRevolution)
		// Mars: orbital period 1.8808158 Earth years
	case "Mars":
		return seconds / (1.8808158 * earthRevolution)
		// Jupiter: orbital period 11.862615 Earth years
	case "Jupiter":
		return seconds / (11.862615 * earthRevolution)
		// Saturn: orbital period 29.447498 Earth years
	case "Saturn":
		return seconds / (29.447498 * earthRevolution)
		// Uranus: orbital period 84.016846 Earth years
	case "Uranus":
		return seconds / (84.016846 * earthRevolution)
		// Neptune: orbital period 164.79132 Earth years
	case "Neptune":
		return seconds / (164.79132 * earthRevolution)
	default:
		return seconds
	}
}
