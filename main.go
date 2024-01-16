package main

import (
	"github.com/rwirdemann/marketsync/context/cmd"
	"github.com/rwirdemann/marketsync/context/excel"
)

const ean = 0

func main() {
	catalog := excel.NewCatalog("bestand.xlsx")
	defer catalog.Close()
	cmd.RootCmd.AddCommand(cmd.NewUploadCmd(catalog))
	cmd.Execute()
}
