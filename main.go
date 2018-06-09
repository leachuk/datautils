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

		//See https://golang.org/pkg/fmt/ for formatting of output (%#v,%+v etc)
		fmt.Printf("Return address [%v]: %+v\n",i, mappify.AddressGeocode(address))
	}

}
