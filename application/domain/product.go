package domain

type Product struct {
	Ean       string  `json:"ean"`
	Inventory int     `json:"inventory"`
	Price     float64 `json:"price"`
}
