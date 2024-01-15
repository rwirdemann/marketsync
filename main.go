package main

import (
	"fmt"
	"time"
)
import "github.com/xuri/excelize/v2"

const ean = 0

func main() {
	f, err := excelize.OpenFile("bestand.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows("bestand")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, row := range rows {
		if i > 0 {
			updateInmventory(row[ean])
			updateLastSync(f, i, time.Now())
		}
	}
	f.Save()
}

func updateLastSync(f *excelize.File, i int, now time.Time) {
	f.SetCellValue("bestand", fmt.Sprintf("E%d", i+1), now)
}

func updateInmventory(ean string) {
	fmt.Printf("ean: %s\n", ean)
}
