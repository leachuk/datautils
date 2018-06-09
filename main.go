package main

import (
	"fmt"
	"github.com/leachuk/datautils/csvservice"
	"github.com/leachuk/datautils/mappify"
	"strconv"
)



func main() {
	csv := csvservice.ReadCSV("/csvservice/datasample.csv")
	var data = [][]string{{"confidence", "streetAddress", "state", "lat", "long"}}

	for i, row := range csv{
		var address = mappify.Address{
			row.Data["streetAddress"],
			row.Data["suburb"],
			row.Data["postCode"],
			row.Data["state"]}

		//See https://golang.org/pkg/fmt/ for formatting of output (%#v,%+v etc)
		var geocodeResponse = mappify.AddressGeocode(address)
		fmt.Printf("Return address [%v]: %+v\n",i, geocodeResponse)

		confidence := strconv.FormatFloat(geocodeResponse.Confidence, 'f', -1, 32)
		streetAddress := geocodeResponse.Result.StreetAddress
		state := geocodeResponse.Result.State
		latitude := strconv.FormatFloat(geocodeResponse.Result.Location.Lat, 'f', -1, 32)
		longitude := strconv.FormatFloat(geocodeResponse.Result.Location.Long, 'f', -1, 32)

		data = append(data, []string{confidence, streetAddress, state, latitude, longitude})
	}

	csvservice.WriteCSV("output.csv", data)

}
