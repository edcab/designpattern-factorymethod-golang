package factory

import (
	"designpattern-factorymethod-golang/product"
	concrete "designpattern-factorymethod-golang/product/impl"
	"fmt"
)

func GetTransportType(quPassengers int) (product.ITransport, error) {
	if quPassengers >= 1 && quPassengers <=5 {
		return  concrete.NewTaxi(), nil
	}
	if quPassengers >= 6 && quPassengers <=15 {
		return  concrete.NewVan(), nil
	}
	if quPassengers >=16 {
		return  concrete.NewBus(), nil
	}
	return nil, fmt.Errorf("wrong quantity of passengers")
}
