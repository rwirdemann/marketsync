package main

import (
	"github.com/joho/godotenv"
	"github.com/rwirdemann/marketsync/context/cmd"
	"github.com/rwirdemann/marketsync/context/config"
	"github.com/rwirdemann/marketsync/context/excel"
	"github.com/rwirdemann/marketsync/context/http"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.Yml{}.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	catalog := excel.NewCatalog("bestand.xlsx", cfg)
	defer catalog.Close()
	marketplace := &http.Marketplace{}
	cmd.RootCmd.AddCommand(cmd.NewUploadCmd(catalog, marketplace))
	cmd.Execute()
}
