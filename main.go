package main

import (
	"fmt"
	"github.com/leachuk/datautils/csvservice"
	"github.com/leachuk/datautils/mappify"
)



func main() {
	csv := csvservice.ReadCSV("/csvservice/datasample.csv")

	for i, row := range csv{
		var address = mappify.Address{
			row.Data["streetAddress"],
			row.Data["suburb"],
			row.Data["postCode"],
			row.Data["state"]}

		fmt.Printf("Return address [%v]: %s\n",i, mappify.AddressGeocode(address))
	}

}
