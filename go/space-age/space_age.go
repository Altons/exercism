// Package space provides age calculation for solar planets
package space

// Planet holds solar system planet names
type Planet string

// Age calculates how old someone is (in years) in Planet X given a number of secods
func Age(s float64, p Planet) float64 {
	var r float64 = s / 31557600
	switch planet := p; planet {
	case "Mercury":
		r = r / 0.2408467
	case "Mars":
		r = r / 1.8808158
	case "Venus":
		r = r / 0.61519726
	case "Saturn":
		r = r / 29.447498
	case "Jupiter":
		r = r / 11.862615
	case "Uranus":
		r = r / 84.016846
	case "Neptune":
		r = r / 164.79132
	default:
	}
	return r
}
