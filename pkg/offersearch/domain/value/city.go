package value

type City struct {
	name        string
	countryCode string
}

// NewCity is intentionally the only way to create a city.
// Value objects are immutable so the struct fields are all unexported.
func NewCity(name string, countryCode string) City {
	return City{
		name:        name,
		countryCode: countryCode,
	}
}

func (c City) Name() string {
	return c.name
}

func (c City) CountryCode() string {
	return c.countryCode
}
