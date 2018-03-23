package main

import (
	"fmt"
	"github.com/leachuk/datautils/mappify"
	"github.com/leachuk/datautils/csvservice"
)

func main() {
	csv := csvservice.ReadCSV("/csvservice/datasample.csv")
	csv2 := csvservice.ReadCSV("/csvservice/datasample2.csv")
	fmt.Printf("CVS1 Header data:%s\n", csv.Heading)
	fmt.Printf("CSV1 Row data:%s\n", csv.Row)
	fmt.Printf("CSV2 Header data:%s\n", csv2.Heading)
	fmt.Printf("CSV2 Row data:%s\n", csv2.Row)
	address := mappify.Address{
		"252 Botany St",
		"Sydney",
		"2000",
		"NSW"}

	fmt.Println("Return data: " + mappify.AddressGeocode(address))
}
