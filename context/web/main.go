package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/rwirdemann/marketsync/application/usecases"
	"github.com/rwirdemann/marketsync/context/config"
	"github.com/rwirdemann/marketsync/context/excel"
	http2 "github.com/rwirdemann/marketsync/context/http"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/upload", uploadHandler)

	port := 4000
	log.Printf("listing on port %d...", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), router)

}

type IndexData struct {
	LastUpload string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := IndexData{
		LastUpload: time.Now().Format(time.DateTime),
	}
	box := packr.NewBox("./templates")
	s, err := box.FindString("index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl, _ := template.New("index").Parse(s)
	tmpl.Execute(w, data)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.Yml{}.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	catalog := excel.NewCatalog("bestand.xlsx", cfg)
	defer catalog.Close()

	marketplace := http2.Marketplace{}
	usecases.Upload(catalog, marketplace)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
