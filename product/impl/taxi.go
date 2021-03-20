package product

import (
	"designpattern-factorymethod-golang/product"
	"designpattern-factorymethod-golang/product/impl/model"
)

type taxi struct {
	model.Transport
}

func NewTaxi() product.ITransport {
	return &taxi{
		Transport: model.Transport{
			Name: "TAXI",
		},
	}
}
