package main

import (
	"fmt"
	"github.com/leachuk/datautils/csvservice"
	"github.com/leachuk/datautils/mappify"
)

func main() {
	csv := csvservice.ReadCSV("/csvservice/datasample.csv")
	fmt.Printf("\nCSV:%s\n", csv)

	address1 := mappify.Address{
		csv[0].Data["streetAddress"],
		csv[0].Data["suburb"],
		csv[0].Data["postCode"],
		csv[0].Data["state"]}

	address2 := mappify.Address{
		csv[1].Data["streetAddress"],
		csv[1].Data["suburb"],
		csv[1].Data["postCode"],
		csv[1].Data["state"]}

	fmt.Println("Return address 1: " + mappify.AddressGeocode(address1))
	fmt.Println("Return address 2: " + mappify.AddressGeocode(address2))
}
