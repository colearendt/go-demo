package main

import (
	"flag"
	"github.com/gorilla/mux"
	"io/fs"
	"log"
	"net/http"
	"time"
)

func main() {
	var dir string
	var useEmbedded bool

	flag.StringVar(&dir, "dir", "./public/", "the directory to serve files from. Defaults to the current dir")
	flag.BoolVar(&useEmbedded, "embedded", hasEmbedded, "use embedded files instead of reading from disk")
	flag.Parse()

	if useEmbedded && !hasEmbedded {
		log.Fatal("Embedded files requested but not available. Build with -tags embed to enable.")
	}
	r := mux.NewRouter()

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
