package excel

import (
	"fmt"
	"github.com/rwirdemann/marketsync/ports/out"
	"github.com/xuri/excelize/v2"
	"log"
	"time"
)

type Catalog struct {
	f    *excelize.File
	rows [][]string
	idx  *int
	cfg  out.Config
}

func NewCatalog(path string, cfg out.Config) *Catalog {
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := f.GetRows(cfg.Source.Sheet)
	if err != nil {
		log.Fatal(err)
	}

	c := &Catalog{f: f, rows: rows}
	c.idx = new(int)
	*c.idx = cfg.Source.FirstContentRow
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
	err := c.f.Save()
	if err != nil {
		log.Fatal(err)
	}
	if err := c.f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (c Catalog) MarkSynced(row []string, t time.Time) {
	err := c.f.SetCellValue("bestand", fmt.Sprintf("E%d", *c.idx), t)
	if err != nil {
		log.Fatal(err)
	}
}
