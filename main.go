package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"crg.eti.br/go/config"
	_ "crg.eti.br/go/config/ini"
	"github.com/google/uuid"
)

type Config struct {
	Port     int    `json:"port" ini:"port" cfg:"port" cfgDefault:"9922"`
	DumpFile string `json:"dump_file" ini:"dump_file" cfg:"dump_file" cfgDefault:"dump.txt"`
}

func saveDumpToFile(dump []byte) {
	f, err := os.OpenFile("dump.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	t := time.Now()
	uuid := uuid.New()

	boundaryStart := fmt.Sprintf("\n- START:%s %s -\n",
		t.Format("2006-01-02 15:04:05"), uuid.String())

	boundaryEnd := fmt.Sprintf("\n- END:%s %s -\n",
		t.Format("2006-01-02 15:04:05"), uuid.String())

	_, err = f.Write([]byte(boundaryStart))
	if err != nil {
		log.Println(err)
		return
	}

	// replace all \r\n to \n
	s := string(dump)
	s = strings.ReplaceAll(s, "\r\n", "\n")
	dump = []byte(s)

	_, err = f.Write(dump)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = f.Write([]byte(boundaryEnd))
	if err != nil {
		log.Println(err)
		return
	}

	err = f.Sync()
	if err != nil {
		log.Println(err)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	saveDumpToFile(reqDump)
	w.Write(reqDump)
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
