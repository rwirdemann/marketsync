package http

import (
	"github.com/rwirdemann/marketsync/application/domain"
	"log"
)

type Marketplace struct{}

func (Marketplace) CreateOrUpdateProduct(product domain.Product) {
	log.Printf("POST /v3/products %v", product)
}
