package main

import (
	"fmt"
	"github.com/leachuk/datautils/csvservice"
	"github.com/leachuk/datautils/mappify"
	"strconv"
	"flag"
)

func main() {
	isDebug := flag.Bool("debug", false, "Enable debug output")
	filename := flag.String("csvfile", "", "CSV filename. Defaults to an empty string")
	flag.Parse()

	if *filename == "" {
		fmt.Println("Error: csv filename argument required.")
	}else{
		csvFilenameArg := *filename
		fmt.Printf("Parsing file: %s\n", csvFilenameArg)
		csv := csvservice.ReadCSV(csvFilenameArg)

		//csv file column headings
		var csvOutput = [][]string{{"confidence", "streetAddress", "state", "lat", "long"}}

		for i, row := range csv{
		var address = mappify.Address{
		row.Data["streetAddress"],
		row.Data["suburb"],
		row.Data["postCode"],
		row.Data["state"]}

		//See https://golang.org/pkg/fmt/ for formatting of output (%#v,%+v etc)
		var geocodeResponse = mappify.AddressGeocode(address)
		if *isDebug {
			fmt.Printf("Return address [%v]: %+v\n", i, geocodeResponse)
		}
		confidence := strconv.FormatFloat(geocodeResponse.Confidence, 'f', -1, 32)
		streetAddress := geocodeResponse.Result.StreetAddress
		state := geocodeResponse.Result.State
		latitude := strconv.FormatFloat(geocodeResponse.Result.Location.Lat, 'f', -1, 32)
		longitude := strconv.FormatFloat(geocodeResponse.Result.Location.Long, 'f', -1, 32)

		csvOutput = append(csvOutput, []string{confidence, streetAddress, state, latitude, longitude})
	}

		csvservice.WriteCSV("output.csv", csvOutput)
		fmt.Println("Successfully created output.csv")
	}

}
