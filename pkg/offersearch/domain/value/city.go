package value

type City struct {
	name        string
	countryCode string
}

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
