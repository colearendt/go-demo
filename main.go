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
	var addr string

	flag.StringVar(&dir, "dir", "./public/", "the directory to serve files from. Defaults to the current dir")
	flag.BoolVar(&useEmbedded, "embedded", hasEmbedded, "use embedded files instead of reading from disk")
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.StringVar(&addr, "addr", "127.0.0.1:8000", "address to listen on (host:port)")
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

	// API endpoint for chart data
	r.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := map[string]interface{}{
			"labels": []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
			"datasets": []map[string]interface{}{
				{
					"label":           "Food & Beverages",
					"data":            []int{100, 102, 105, 108, 110, 113, 115, 118, 120, 121, 123, 125},
					"borderColor":     "#667eea",
					"backgroundColor": "rgba(102, 126, 234, 0.1)",
				},
				{
					"label":           "Personal Care",
					"data":            []int{100, 101, 102, 103, 105, 106, 108, 109, 110, 111, 112, 114},
					"borderColor":     "#764ba2",
					"backgroundColor": "rgba(118, 75, 162, 0.1)",
				},
				{
					"label":           "Household Items",
					"data":            []int{100, 101, 103, 104, 106, 107, 108, 109, 110, 111, 111, 112},
					"borderColor":     "#f093fb",
					"backgroundColor": "rgba(240, 147, 251, 0.1)",
				},
				{
					"label":           "Electronics",
					"data":            []int{100, 99, 98, 97, 97, 96, 96, 95, 95, 94, 94, 93},
					"borderColor":     "#4facfe",
					"backgroundColor": "rgba(79, 172, 254, 0.1)",
				},
				{
					"label":           "Clothing",
					"data":            []int{100, 100, 101, 102, 103, 104, 105, 106, 107, 108, 108, 109},
					"borderColor":     "#43e97b",
					"backgroundColor": "rgba(67, 233, 123, 0.1)",
				},
				{
					"label":           "Transportation",
					"data":            []int{100, 103, 106, 109, 112, 114, 116, 117, 118, 119, 120, 122},
					"borderColor":     "#fa709a",
					"backgroundColor": "rgba(250, 112, 154, 0.1)",
				},
			},
		}
		json.NewEncoder(w).Encode(data)
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
		Addr:    addr,
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
