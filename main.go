package main

import (
	"designpattern-factorymethod-golang/product"
	"designpattern-factorymethod-golang/product/factory"
	"fmt"
)

func main() {
	taxi, _ := factory.GetTransportType(1)
	van, _ := factory.GetTransportType(6)
	bus, _ := factory.GetTransportType(20)

	printDetails(taxi)
	printDetails(van)
	printDetails(bus)
}

func printDetails(g product.ITransport) {
	fmt.Printf("Transport assigned: %s", g.GetName())
	fmt.Println()
}
