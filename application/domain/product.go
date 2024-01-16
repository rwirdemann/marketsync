package domain

import "fmt"

type Product struct {
	Ean       string  `json:"ean"`
	Inventory int     `json:"inventory"`
	Price     float64 `json:"price"`
}

func (p Product) String() string {
	return fmt.Sprintf("{ean: %s}", p.Ean)
}
