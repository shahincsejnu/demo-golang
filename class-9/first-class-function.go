package main

// import "fmt"

// type sum func(a int, b int) int

// func temp(s func(a int, b int) int) int {
// 	res := s(20, 30)
// 	return res
// }

// func temp2() func(a int, b int) int {
// 	f := func(a int, b int) int {
// 		return a + b
// 	}

// 	return f
// }

// func main() {
// 	val := func(x int) int {
// 		fmt.Println("Hello World!", x)
// 		return x
// 	}

// 	fmt.Println(val(20))

// 	//var arr [5]int = [5]int{1, 2, 3, 4, 5}

// 	var sumVar sum
// 	sumVar = func(a, b int) int {
// 		return a + b
// 	}

// 	fmt.Println(sumVar(10, 20))

// 	s := func(a int, b int) int {
// 		return a + b
// 	}

// 	fmt.Println(temp(s))

// 	ss := temp2()

// 	fmt.Println(temp(ss))
// }
