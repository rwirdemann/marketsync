package usecases

import (
	"github.com/rwirdemann/marketsync/application/domain"
	"github.com/rwirdemann/marketsync/ports/out"
	"time"
)

const ean = 0

// Upload uploads the catalog to the marketplace.
func Upload(catalog out.Catalog, marketplace out.Marketplace) error {
	if err := marketplace.Login(); err != nil {
		return err
	}

	t := time.Now()
	for hasNext := true; hasNext; hasNext = catalog.HasNext() {
		row := catalog.NextRow()
		process(row, marketplace)
		catalog.MarkSynced(row, t)
	}
	return nil
}

func process(row []string, marketplace out.Marketplace) {
	p := domain.Product{
		ProductReference: "UBN-11779",
		Sku:              "3858389911565",
		Ean:              "3858389911565",
		Pzn:              "PZN-4908802",
		Mpn:              "H2G2-42",
		Moin:             "M00A1234BC",
		ReleaseDate:      time.Now(),
		ProductDescription: domain.ProductDescription{
			Category:       "Outdoorjacke",
			BrandId:        "CMWBBRW2",
			ProductLine:    "501",
			Manufacturer:   "3M",
			ProductionDate: time.Now(),
			MultiPack:      true,
			Bundle:         false,
			FscCertified:   true,
			Disposal:       false,
			ProductUrl:     "http://myproduct.somewhere.com/productname/",
			Description:    "<p>Some example words...<b>in bold</b>...some more</p>",
			BulletPoints: []string{
				"My top key information...",
			},
			Attributes: []domain.Attribute{
				{
					Name:       "Bundweite",
					Values:     []string{"34"},
					Additional: false,
				},
			},
		},
		MediaAssets: []domain.MediaAsset{
			{
				Type:     "IMAGE",
				Location: "http://apartners.url/image-location",
			},
		},
		Order: domain.Order{MaxOrderQuantity: domain.MaxOrderQuantity{
			Quantity:     10,
			PeriodInDays: 7,
		}},
		Delivery: domain.Delivery{
			Type:         "PARCEL",
			DeliveryTime: 1,
		},
		Pricing: domain.Pricing{
			StandardPrice: domain.Price{
				Amount:   19.95,
				Currency: "EUR",
			},
			Vat: "FULL",
			Msrp: domain.Price{
				Amount:   19.95,
				Currency: "EUR",
			},
			Sale: domain.Sale{
				SalePrice: domain.Price{
					Amount:   19.95,
					Currency: "EUR",
				},
				StartDate: time.Now(),
				EndDate:   time.Now(),
			},
			NormPriceInfo: domain.NormPriceInfo{
				NormAmount:  100,
				NormUnit:    "g",
				SalesAmount: 500,
				SalesUnit:   "g",
			},
		},
		Logistics: domain.Logistics{
			PackingUnitCount: 3,
			PackingUnits: []domain.PackingUnit{
				{
					Weight: 365,
					Width:  600,
					Height: 200,
					Length: 300,
				},
			},
		},
	}

	marketplace.CreateOrUpdateProduct(p)
}
