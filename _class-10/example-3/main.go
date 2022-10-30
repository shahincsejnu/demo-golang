package main

import (
	"embed"
	"net/http"
)

//go:embed static
var contents embed.FS

func main() {
	r := http.NewServeMux()

	r.Handle("/", http.FileServer(http.Dir("static")))

	http.ListenAndServe(":8080", r)

	
}
