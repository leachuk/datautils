package mappify

import "fmt"

type Address struct {
	StreetAddress	string
	Suburb	string
	Postcode	string
	State	string
}

//provide custom string formatter to map to Mappify requirement. Will replace with JSON later.
//{"streetAddress":"252 Botany St","suburb":"Sydney","postcode":2000,"state":"NSW"}
func (a Address) String() string {
	return fmt.Sprintf(
		"\"streetAddress\":\"%s\", " +
		"\"suburb\":\"%s\", " +
		"\"postCode\":\"%s\", " +
		"\"state\":\"%s\"", a.StreetAddress, a.Suburb, a.Postcode, a.State)
}

func AddressGeocode(address Address) string {
	return "Mappify this: " + fmt.Sprintln(address)
}

