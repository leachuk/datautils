package main

import (
	"fmt"
	//"github.com/leachuk/datautils/mappify"
	"github.com/leachuk/datautils/csvservice"
	//"github.com/leachuk/datautils/mappify"
	//"encoding/json"
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
	fmt.Println(csv[0].Data["postCode"])

	//address1 := mappify.Address{
	//	csv[0].Data[0],
	//	"Sydney",
	//	"2000",
	//	"NSW"}
	//addressJson, _ := json.Marshal(address1)
	//fmt.Println(string(addressJson))
	//fmt.Println("Return data: " + mappify.AddressGeocode(address))
}
