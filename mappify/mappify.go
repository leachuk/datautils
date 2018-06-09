package mappify

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
	"encoding/json"
)

const MappifyHost = "http://mappify.io/api/rpc"
const MappifyAddrGeocodeEndpoint = MappifyHost + "/address/geocode"

type Address struct {
	StreetAddress	string	`json:"streetAddress"`
	Suburb			string	`json:"suburb"`
	Postcode		string	`json:"postCode"`
	State			string	`json:"state"`
}

//note. All types must be set correctly (i.e. float64/string/int) or things are intermittently incorrect
//e.g. Lat & Long were string, which caused StreetAdress to be empty, while State was correctly set.
type AddressGeocodeResponse struct {
	Result struct {
		Location struct{
			Lat		float64	`json:"lat"`
			Long 	float64	`json:"lon"`
		} `json:"location"`
		StreetAddress	string	`json:"streetAddress"`
		State			string	`json:"state"`
		Suburb			string	`json:"suburb"`
	} `json:"result"`
	Confidence	float64
	Type		string
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

func AddressGeocode(address Address) AddressGeocodeResponse {
	//jsonData := fmt.Sprintf("{%s}",address)
	jsonData, _ := json.Marshal(address)
	fmt.Println("Input data:" + string(jsonData))

	client := &http.Client{}
	req, err := http.NewRequest("POST",
		MappifyAddrGeocodeEndpoint,
		bytes.NewBufferString(string(jsonData)))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http POST error: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var jsonResponse AddressGeocodeResponse
	err = json.Unmarshal([]byte(body), &jsonResponse)
	if err != nil {
		fmt.Println("error:", err)
	}

	return AddressGeocodeResponse(jsonResponse)
}

