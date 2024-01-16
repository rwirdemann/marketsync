package main

import (
	"github.com/rwirdemann/marketsync/context/cmd"
	"github.com/rwirdemann/marketsync/context/config"
	"github.com/rwirdemann/marketsync/context/excel"
	"github.com/rwirdemann/marketsync/context/http"
	"log"
)

const ean = 0

func main() {
	cfg, err := config.Yml{}.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	catalog := excel.NewCatalog("bestand.xlsx", cfg)
	defer catalog.Close()
	marketplace := http.Marketplace{}
	cmd.RootCmd.AddCommand(cmd.NewUploadCmd(catalog, marketplace))
	cmd.Execute()
}
