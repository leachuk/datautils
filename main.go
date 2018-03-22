package main

import (
	"fmt"
	"github.com/leachuk/datautils/mappify"
)

func main() {
	address := mappify.Address{
		"252 Botany St",
		"Sydney",
		"2000",
		"NSW"}

	fmt.Println(mappify.AddressGeocode(address))
}
