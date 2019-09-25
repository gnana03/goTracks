package space_age

func Age(seconds float64,planet string) float64{

	switch planet {
	case "Earth":
		return seconds/(365.25*24*60*60)
	case "Mercury":
		return seconds/(0.2408467*365.25*24*60*60)
	case "Venus":
		return seconds/(0.61519726*365.25*24*60*60)
	case "Mars":
		return seconds/(1.8808158*365.25*24*60*60)
	case "Jupiter":
		return seconds/(11.862615*365.25*24*60*60)
	case "Saturn":
		return seconds/(29.447498*365.25*24*60*60)
	case "Uranus":
		return seconds/(84.016846*365.25*24*60*60)
	case "Neptune":
		return seconds/(164.79132*365.25*24*60*60)
	}
	return 0.0
}

