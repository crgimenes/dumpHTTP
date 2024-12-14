package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	ANSI_RESET         = "\x1b[0m"
	ANSI_WHITE         = "\x1b[37m"
	ANSI_BRIGHT_YELLOW = "\x1b[93m"
	ANSI_BRIGHT_CYAN   = "\x1b[96m"
)

var (
	listenAddr *string
	dumpFile   *string
	mx         sync.Mutex
)

func saveDumpToFile(dump []byte) {
	f, err := os.OpenFile(
		*dumpFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	t := time.Now()
	id := uuid.New()

	boundary := fmt.Sprintf(
		"\n- boundary-%s-%s\n",
		t.Format("2006-01-02 15:04:05"),
		id.String())

	_, err = f.Write([]byte(boundary))
	if err != nil {
		log.Println(err)
		return
	}

	_, err = f.Write(dump)
	if err != nil {
		log.Println(err)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mx.Lock()
	defer mx.Unlock()

	fmt.Printf("%sRequest:\n%s%s%s----------------------------------------%s\n",
		ANSI_BRIGHT_YELLOW,
		ANSI_WHITE,
		string(reqDump),
		ANSI_BRIGHT_CYAN,
		ANSI_RESET)

	saveDumpToFile(reqDump)
	w.Write(reqDump)
}

func main() {
	listenAddr = flag.String("listenAddr", ":8080", "Address to listen")
	dumpFile = flag.String("dumpFile", "dump.txt", "File to save dump")
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	http.HandleFunc("/", handler)
	log.Printf("Starting dumpHTTP server on %s\n", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
