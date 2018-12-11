// hexe exposes an arbitrary binary via simplified HTTP.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var (
	listen = flag.String("l", ":8080", "HTTP listening address")
)

func handler(w http.ResponseWriter, r *http.Request) {
	args := flag.Args()
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = w
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(w, "error: %v", err)
		log.Printf("error: %v", err)
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