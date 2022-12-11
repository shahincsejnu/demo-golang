package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type response struct {
// 	OK          bool
// 	ServiceName string
// }

// func main() {
// 	jsn1, err := json.Marshal(true)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println(string(jsn1))

// 	jsn2, err := json.Marshal(1)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println(string(jsn2))

// 	res1 := &response{
// 		OK:          true,
// 		ServiceName: "demo-service",
// 	}

// 	jsn3, err := json.MarshalIndent(res1, "", "  ")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println(string(jsn3))

// 	var res2 response
// 	fmt.Println(res2)

// 	err = json.Unmarshal(jsn3, &res2)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println(res2)

// 	err = json.NewEncoder(os.Stdout).Encode(res1)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	var res3 response

// 	err = json.NewDecoder(os.Stdin).Decode(&res3)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	fmt.Println(res3)
// }
