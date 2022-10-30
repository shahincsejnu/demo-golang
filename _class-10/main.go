package main

// import (
// 	"embed"
// 	"log"
// 	"net/http"
// )

// //go:embed *
// var content embed.FS

// func main() {
// 	mutex := http.NewServeMux()
// 	mutex.Handle("/", http.FileServer(http.FS(content)))
// 	err := http.ListenAndServe(":8080", mutex)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

import (
	_ "embed"
	"fmt"
)

func main() {
	fmt.Printf("Version %q\n", version)
}
