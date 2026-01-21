package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

var version = "dev"

func main() {
	var dir string
	var useEmbedded bool
	var showVersion bool

	flag.StringVar(&dir, "dir", "./public/", "the directory to serve files from. Defaults to the current dir")
	flag.BoolVar(&useEmbedded, "embedded", hasEmbedded, "use embedded files instead of reading from disk")
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.Parse()

	if showVersion {
		fmt.Printf("go-demo version %s\n", version)
		os.Exit(0)
	}

	if useEmbedded && !hasEmbedded {
		log.Fatal("Embedded files requested but not available. Build with -tags embed to enable.")
	}
	r := mux.NewRouter()

	// API endpoint for version
	r.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("format") == "plain" {
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintln(w, version)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"version": version,
			})
		}
	}).Methods("GET")

	// This will serve files under http://localhost:8000/<filename>
	var handler http.Handler
	if useEmbedded {
		subFS, err := fs.Sub(embeddedFiles, "public")
		if err != nil {
			log.Fatal(err)
		}
		handler = http.FileServer(http.FS(subFS))
	} else {
		handler = http.FileServer(http.Dir(dir))
	}
	r.PathPrefix("/").Handler(handler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if useEmbedded {
		log.Printf("Serving embedded files on http://%s\n", srv.Addr)
	} else {
		log.Printf("Serving %s on http://%s\n", dir, srv.Addr)
	}
	log.Fatal(srv.ListenAndServe())
}
