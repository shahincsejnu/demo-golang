package main

import "fmt"

// // type TwoSum struct {
// // 	X int
// // 	Y int
// // }

// // type ThSum struct {
// // 	X int
// // 	Y int
// // 	Z int
// // }

// // type FileSystem interface {
// // 	Read() []byte
// // 	Write(data []byte)
// // 	List(path string) int
// // }

// // type AbstractionForSummation interface {
// // 	Sum() int
// // }

// // func (t TwoSum) Sum() int {
// // 	return t.X + t.Y
// // }

// // func (th ThSum) Sum() int {
// // 	return th.X + th.Y + th.Z
// // }

// // type MyError struct {
// // 	Str string
// // }

// // func (m *MyError) Error() string {
// // 	return "oka"
// // }

// // func helper(anyType AbstractionForSummation) (int, error) {
// // 	return anyType.Sum(), errors.New("this is a dummy error")
// // }

// func main() {
// 	// t := TwoSum{10, 20}

// 	// th := ThSum{10, 20, 30}

// 	// var col = []AbstractionForSummation{t, th}

// 	// for _, val := range col {
// 	// 	res, err := helper(val)
// 	// 	if err != nil {
// 	// 		fmt.Println(res)
// 	// 		fmt.Println(err.Error())
// 	// 		return
// 	// 	}

// 	// 	// val (type, value)

// 	// 	fmt.Printf("type : %T, result: %v", val, val)
// 	// 	fmt.Println()
// 	// 	fmt.Println(res)
// 	// }

// 	// fmt.Println("\n\n\n")
// 	// var temp interface{}
// 	// fmt.Println(temp)
// 	// fmt.Printf("type: %T, value: %v", temp, temp)
// 	// fmt.Println()

// 	// temp = 10
// 	// fmt.Println(temp)
// 	// fmt.Printf("type: %T, value: %v", temp, temp)
// 	// fmt.Println()

// 	// temp = "string"
// 	// fmt.Println(temp)
// 	// fmt.Printf("type: %T, value: %v", temp, temp)
// 	// fmt.Println()

// 	// //var mp map[string]interface{}

// 	// var ab AbstractionForSummation
// 	// fmt.Printf("type: %T, value: %v", ab, ab)
// 	// ab.Sum()
// 	// fmt.Println()

// 	// var temp interface{}
// 	// fmt.Printf("%T %v", temp, temp)
// 	// if temp == nil {
// 	// 	fmt.Println("it's nil")
// 	// } else {
// 	// 	fmt.Println("not nil")
// 	// }
// 	// var ts *TwoSum

// 	// temp = ts

// 	// fmt.Println(ts)

// 	// fmt.Printf("%T %v", temp, temp)
// 	// fmt.Println()

// 	// if temp == nil {
// 	// 	fmt.Println("it's nil")
// 	// } else {
// 	// 	fmt.Println("not nil")
// 	// }

// 	// if ts == nil {
// 	// 	fmt.Println("it's nil")
// 	// } else {
// 	// 	fmt.Println("not nil")
// 	// }

// 	// var temp interface{}
// 	// temp = 10

// 	// //var val int

// 	// val, ok := temp.(int)

// 	// Temp(val)

// 	// temp(1)
// 	// temp("one")
// 	// temp(true)
// 	// temp(12.5)

// }

// // func temp(val any) {
// // 	switch val.(type) {
// // 	case int:
// // 		fmt.Println("integer")
// // 	case string:
// // 		fmt.Println("string")
// // 	case bool:
// // 		fmt.Println("bool")
// // 	default:
// // 		fmt.Println("not supported type")
// // 	}
// // }

func main() {
	val := func(x int) int {
		return x
	}

	fmt.Println(val(10))
}
