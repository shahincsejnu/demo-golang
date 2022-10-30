package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./text2.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	f.WriteString("hello, this is text2")

	fmt.Fprintln(f, "new line")
}
