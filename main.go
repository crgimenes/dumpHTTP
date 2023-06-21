package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"crg.eti.br/go/config"
	_ "crg.eti.br/go/config/ini"
)

type Config struct {
	Port int `json:"port" ini:"port" cfg:"port" cfgDefault:"9922"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s", string(reqDump))

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{"status": "ok"}`))
}

func main() {
	cfg := Config{}

	config.File = "config.ini"
	err := config.Parse(&cfg)
	if err != nil {
		println(err)
		return
	}

	http.HandleFunc("/", indexHandler)
	log.Printf("Starting dumpHTTP server at port: %d\n", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}
