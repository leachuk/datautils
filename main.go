package main

import (
	"fmt"
	"github.com/leachuk/datautils/csvservice"
	"github.com/leachuk/datautils/mappify"
	"strconv"
	"os"
)



func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Error: csv filename argument required.")
	}else{
		csvFilenameArg := os.Args[1]
		fmt.Printf("Parsing file: %s\n", csvFilenameArg)
		csv := csvservice.ReadCSV(csvFilenameArg)

		var csvOutput = [][]string{{"confidence", "streetAddress", "state", "lat", "long"}}

		for i, row := range csv{
		var address = mappify.Address{
		row.Data["streetAddress"],
		row.Data["suburb"],
		row.Data["postCode"],
		row.Data["state"]}

		//See https://golang.org/pkg/fmt/ for formatting of output (%#v,%+v etc)
		var geocodeResponse = mappify.AddressGeocode(address)
		fmt.Printf("Return address [%v]: %+v\n", i, geocodeResponse)

		confidence := strconv.FormatFloat(geocodeResponse.Confidence, 'f', -1, 32)
		streetAddress := geocodeResponse.Result.StreetAddress
		state := geocodeResponse.Result.State
		latitude := strconv.FormatFloat(geocodeResponse.Result.Location.Lat, 'f', -1, 32)
		longitude := strconv.FormatFloat(geocodeResponse.Result.Location.Long, 'f', -1, 32)

		csvOutput = append(csvOutput, []string{confidence, streetAddress, state, latitude, longitude})
	}

		csvservice.WriteCSV("output.csv", csvOutput)
	}

}
