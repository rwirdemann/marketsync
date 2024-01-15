package main

import (
	"fmt"
	"github.com/rwirdemann/marketsync/application/domain"
	"github.com/rwirdemann/marketsync/application/usecases"
	"github.com/rwirdemann/marketsync/context/excel"
	"time"
)
import "github.com/xuri/excelize/v2"

const ean = 0

func main() {
	catalog := excel.NewCatalog("bestand.xlsx")
	defer catalog.Close()
	usecases.Upload(catalog)

	// f.Save()
}

func updateLastSync(f *excelize.File, i int, now time.Time) {
	f.SetCellValue("bestand", fmt.Sprintf("E%d", i+1), now)
}

func sync(p domain.Product) {
	fmt.Printf("ean: %s\n", p.Ean)
}
