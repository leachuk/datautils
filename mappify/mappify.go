package mappify

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

const MappifyHost = "http://mappify.io/api/rpc"
const MappifyAddrGeocodeEndpoint = MappifyHost + "/address/geocode"

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
	data := fmt.Sprintf("{%s}",address)
	fmt.Println("Input data:" + data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", MappifyAddrGeocodeEndpoint,
		bytes.NewBufferString(data))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http POST error: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

