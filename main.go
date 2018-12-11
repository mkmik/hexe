// hexe exposes an arbitrary binary via simplified HTTP.
package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var (
	listen = flag.String("l", ":8080", "HTTP listening address")
)

func handler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	args := flag.Args()
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err := io.Copy(w, &buf)
	if err != nil {
		log.Printf("error writing to client: %v", err)
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)

	log.Printf("HTTP requests %q will execute %q", *listen, flag.Args())
	if err := http.ListenAndServe(*listen, nil); err != nil {
		log.Fatal(err)
	}
}