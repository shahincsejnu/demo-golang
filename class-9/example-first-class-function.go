package main

import (
	"fmt"
	"sort"
)

func main() {
	temp := []struct {
		Name string
		Id   int
	}{
		{
			"Third",
			3,
		},
		{
			"Second",
			2,
		},
		{
			"First",
			1,
		},
	}

	for key, val := range temp {
		fmt.Println(key, val)
	}
	fmt.Println()

	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].Id < temp[j].Id
	})

	for key, val := range temp {
		fmt.Println(key, val)
	}
}
