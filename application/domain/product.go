package domain

import (
	"fmt"
	"time"
)

type Attribute struct {
	Name       string   `json:"name"`
	Values     []string `json:"values"`
	Additional bool     `json:"additional"`
}

type ProductDescription struct {
	Category       string      `json:"category"`
	BrandId        string      `json:"brandId"`
	ProductLine    string      `json:"productLine"`
	Manufacturer   string      `json:"manufacturer"`
	ProductionDate time.Time   `json:"productionDate"`
	MultiPack      bool        `json:"multiPack"`
	Bundle         bool        `json:"bundle"`
	FscCertified   bool        `json:"fscCertified"`
	Disposal       bool        `json:"disposal"`
	ProductUrl     string      `json:"productUrl"`
	Description    string      `json:"description"`
	BulletPoints   []string    `json:"bulletPoints"`
	Attributes     []Attribute `json:"attributes"`
}

type MediaAsset struct {
	Type     string `json:"type"`
	Location string `json:"location"`
}

type MaxOrderQuantity struct {
	Quantity     int `json:"quantity"`
	PeriodInDays int `json:"periodInDays"`
}

type Order struct {
	MaxOrderQuantity MaxOrderQuantity `json:"maxOrderQuantity"`
}

type Delivery struct {
	Type         string `json:"type"`
	DeliveryTime int    `json:"deliveryTime"`
}

type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type Sale struct {
	SalePrice Price     `json:"salePrice"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type NormPriceInfo struct {
	NormAmount  int    `json:"normAmount"`
	NormUnit    string `json:"normUnit"`
	SalesAmount int    `json:"salesAmount"`
	SalesUnit   string `json:"salesUnit"`
}

type Pricing struct {
	StandardPrice Price         `json:"standardPrice"`
	Vat           string        `json:"vat"`
	Msrp          Price         `json:"msrp"`
	Sale          Sale          `json:"sale"`
	NormPriceInfo NormPriceInfo `json:"normPriceInfo"`
}

type PackingUnit struct {
	Weight int `json:"weight"`
	Width  int `json:"width"`
	Height int `json:"height"`
	Length int `json:"length"`
}

type Logistics struct {
	PackingUnitCount int           `json:"packingUnitCount"`
	PackingUnits     []PackingUnit `json:"packingUnits"`
}

type Product struct {
	ProductReference   string             `json:"productReference"`
	Sku                string             `json:"sku"`
	Ean                string             `json:"ean"`
	Pzn                string             `json:"pzn"`
	Mpn                string             `json:"mpn"`
	Moin               string             `json:"moin"`
	ReleaseDate        time.Time          `json:"releaseDate"`
	ProductDescription ProductDescription `json:"productDescription"`
	MediaAssets        []MediaAsset       `json:"mediaAssets"`
	Order              Order              `json:"order"`
	Delivery           Delivery           `json:"delivery"`
	Pricing            Pricing            `json:"pricing"`
	Logistics          Logistics          `json:"logistics"`
}

func (p Product) String() string {
	return fmt.Sprintf("{ean: %s}", p.Ean)
}
