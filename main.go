package main

import (
	"fmt"
	"github.com/leachuk/datautils/mappify"
)

func main() {
	//todo: parse csv file and extract address
	address := mappify.Address{
		"252 Botany St",
		"Sydney",
		"2000",
		"NSW"}

	fmt.Println("Return data: " + mappify.AddressGeocode(address))
}
