package main

import (
	"embed"
	"fmt"
	"net/http"
	"strings"
)

//go:embed data/words.txt
var data []byte

//go:embed public
var contents embed.FS

func home() http.Handler {
	return http.FileServer(http.FS(contents))
}

func main() {
	fmt.Println(string(data))

	words := strings.Split(string(data), "\n")

	fmt.Println(words)

	r := http.NewServeMux()

	r.Handle("/", home())

	fmt.Println("starting the server...")
	http.ListenAndServe(":8080", r)
}
