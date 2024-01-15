package main

import (
	"github.com/rwirdemann/marketsync/application/usecases"
	"github.com/rwirdemann/marketsync/context/excel"
)

const ean = 0

func main() {
	catalog := excel.NewCatalog("bestand.xlsx")
	defer catalog.Close()
	usecases.Upload(catalog)
}
