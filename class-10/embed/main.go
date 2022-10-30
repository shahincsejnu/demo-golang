package main

import (
	"embed"
	"fmt"
	"net/http"
)

// //go:embed text.txt
// var contents string

// //go:embed text.txt
// var contents2 []byte

//go:embed *
var f embed.FS

func main() {
	// contents, _ := os.ReadFile("./text.txt")

	// fmt.Println(string(contents))

	// fmt.Println(contents)

	// fmt.Println()

	// fmt.Println(string(contents2))

	// contents, err := f.ReadFile("text.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(string(contents))
	// fmt.Println()

	// contents2, err := f.ReadFile("text2.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(string(contents2))

	//f.ReadFile("temp/temp2/text4.txt")

	// rootDir := "temp"

	// dirs, err := f.ReadDir(rootDir)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// for _, dr := range dirs {
	// 	if dr.IsDir() {
	// 		dirs2, err := f.ReadDir(filepath.Join(rootDir, dr.Name()))
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 			return
	// 		}

	// 		fmt.Println(len((dirs2)))
	// 	} else {
	// 		contents, err := f.ReadFile(filepath.Join(rootDir, dr.Name()))
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 			return
	// 		}
	// 		fmt.Println(dr.Name(), "contents are: ", string(contents))
	// 	}
	// }

	r := http.NewServeMux()

	//r.Handle("/", http.FileServer(http.FS(f)))

	r.HandleFunc("/home", Home)

	fmt.Println("server is running...")
	http.ListenAndServe(":8080", r)

	
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}
