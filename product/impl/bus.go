package product

import (
	"designpattern-factorymethod-golang/product"
	"designpattern-factorymethod-golang/product/impl/model"
)

type bus struct {
	model.Transport // indirectly implement all methods of iGun and hence are of iTransport type
}

func NewBus() product.ITransport {
	return &bus{
		Transport: model.Transport{
			Name: "BUS",
		},
	}
}
