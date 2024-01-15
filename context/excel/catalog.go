package excel

import (
	"github.com/xuri/excelize/v2"
	"log"
)

type Catalog struct {
	f    *excelize.File
	rows [][]string
	idx  *int
}

func NewCatalog(path string) *Catalog {
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := f.GetRows("bestand")
	if err != nil {
		log.Fatal(err)
	}

	c := &Catalog{f: f, rows: rows}
	c.idx = new(int)
	*c.idx = 1
	return c
}

func (c Catalog) NextRow() []string {
	r := c.rows[*c.idx]
	*c.idx = *c.idx + 1
	return r
}

func (c Catalog) HasNext() bool {
	return *c.idx < len(c.rows)
}

func (c Catalog) Close() {
	if err := c.f.Close(); err != nil {
		log.Fatal(err)
	}
}
