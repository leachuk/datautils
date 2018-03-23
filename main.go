package main

import (
	"fmt"
	//"github.com/leachuk/datautils/mappify"
	"github.com/leachuk/datautils/csvservice"
)

func main() {
	csv := csvservice.ReadCSV("/csvservice/datasample.csv")
	fmt.Printf("\nCSV:%s\n", csv)

	//var m = make(map[string]string)
	for i := 0; i < len(csv); i++ {
		fmt.Println(csv[i])
		//for _, headingitem := range csv.Heading {
		//	m[headingitem] = csv.Row[i]
		//}
	}
	//fmt.Println(m)

	//address := mappify.Address{
	//	"252 Botany St",
	//	"Sydney",
	//	"2000",
	//	"NSW"}

	//fmt.Println("Return data: " + mappify.AddressGeocode(address))
}
