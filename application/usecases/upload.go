package usecases

import (
	"fmt"
	"github.com/rwirdemann/marketsync/application/domain"
	"github.com/rwirdemann/marketsync/ports/out"
	"time"
)

const ean = 0

func Upload(catalog out.Catalog) {
	t := time.Now()
	for hasNext := true; hasNext; hasNext = catalog.HasNext() {
		row := catalog.NextRow()
		process(row)
		catalog.MarkSynced(row, t)
	}
}

func process(row []string) {
	p := domain.Product{
		Ean: row[ean],
	}
	fmt.Printf("ean: %s\n", p.Ean)
}
