package out

import "github.com/rwirdemann/marketsync/application/domain"

type Marketplace interface {
	CreateOrUpdateProduct(product domain.Product) error
	Login() error
}
