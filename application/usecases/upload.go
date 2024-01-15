package usecases

import (
	"fmt"
	"github.com/rwirdemann/marketsync/application/domain"
	"github.com/rwirdemann/marketsync/ports/out"
)

const ean = 0

func Upload(catalog out.Catalog) {
	for hasNext := true; hasNext; hasNext = catalog.HasNext() {
		process(catalog.NextRow())
	}
}

func process(row []string) {
	p := domain.Product{
		Ean: row[ean],
	}
	fmt.Printf("ean: %s\n", p.Ean)
}
