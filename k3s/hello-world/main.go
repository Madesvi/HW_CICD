package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello World from Pod %s!\n", hostname)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":30757", nil))
}
