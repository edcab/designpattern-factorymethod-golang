package product

import (
	"designpattern-factorymethod-golang/product"
	"designpattern-factorymethod-golang/product/impl/model"
)

type van struct {
	model.Transport
}

func NewVan() product.ITransport {
	return &van{
		Transport: model.Transport{
			Name: "VAN",
		},
	}
}