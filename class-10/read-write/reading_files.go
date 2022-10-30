package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// contents, err := os.ReadFile("./text.txt")
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// 	return
// 	// }

// 	// fmt.Println(string(contents))

// 	f, err := os.Open("./text.txt")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	defer func() {
// 		err := f.Close()
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 	}()

// 	// r := bufio.NewReader(f)
// 	// data := make([]byte, 5)

// 	// for {
// 	// 	_, err := r.Read(data)
// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}

// 	// 	if err != nil {
// 	// 		fmt.Println(err.Error())
// 	// 		return
// 	// 	}

// 	// 	fmt.Println(string(data))
// 	// }

// 	s := bufio.NewScanner(f)

// 	for s.Scan() {
// 		fmt.Println(s.Text())
// 		fmt.Println()
// 	}

// 	err = s.Err()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// }
