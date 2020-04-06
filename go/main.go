package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	// todo: add version
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	addr := fmt.Sprintf("0.0.0.0:%s", port)

	log.Info("running on port ", port)
	http.HandleFunc("/", handler) //each request calls handler
	log.Fatal(http.ListenAndServe(addr, nil))
}

// handler echos the Path component of the request URL r.
// %q  a single-quoted character literal safely escaped with Go syntax.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
