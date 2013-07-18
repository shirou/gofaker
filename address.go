package gofaker


type Address struct {
	Country string
	City string
	State string
	Street string
	Building_number string

	Latitude string
	Longitude string

	Zipcode string
}

func NewAddress(locale string) Address{
	b := Address{}

	b.City = GetValue(locale, "cities")
	b.State = GetValue(locale, "states")
	b.Street = GetValue(locale, "streets")
	b.Building_number = GetValue(locale, "building_numbers")
	b.Zipcode = GetValue(locale, "zipcode")


	return b
}
