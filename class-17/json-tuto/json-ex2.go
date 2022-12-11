package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	OK          bool   `json:"ok"`
	ServiceName string `json:"-"`
}

func main() {
	res1 := &response{
		OK:          true,
		ServiceName: "demo-service",
	}

	jsn3, err := json.MarshalIndent(res1, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(jsn3))
}
