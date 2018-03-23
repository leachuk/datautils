package main

import (
	"fmt"
	"github.com/leachuk/datautils/mappify"
	"github.com/leachuk/datautils/csvservice"
)

func main() {
	//todo: parse csv file and extract address
	csv := csvservice.ReadCSV("/csvservice/datasample.csv")

	fmt.Printf("Header data:%s\n", csv.Heading)
	fmt.Printf("Row data:%s\n", csv.Row)

	address := mappify.Address{
		"252 Botany St",
		"Sydney",
		"2000",
		"NSW"}

	fmt.Println("Return data: " + mappify.AddressGeocode(address))
}
