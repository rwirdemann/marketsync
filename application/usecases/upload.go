package usecases

import (
	"github.com/rwirdemann/marketsync/application/domain"
	"github.com/rwirdemann/marketsync/ports/out"
	"time"
)

const ean = 0

// Upload uploads the catalog to the marketplace.
func Upload(catalog out.Catalog, marketplace out.Marketplace) {
	t := time.Now()
	for hasNext := true; hasNext; hasNext = catalog.HasNext() {
		row := catalog.NextRow()
		process(row, marketplace)
		catalog.MarkSynced(row, t)
	}
}

func process(row []string, marketplace out.Marketplace) {
	p := domain.Product{
		Ean: row[ean],
	}
	marketplace.CreateOrUpdateProduct(p)
}
