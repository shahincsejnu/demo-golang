package main

// type Student struct {
// 	Name string
// 	Id   int
// }

// type Teacher struct {
// 	Name string
// 	Id   int
// }

// func helper(i interface{}) {

// 	// fmt.Println(i)

// 	// reflect.ValueOf(i).Elem().FieldByName("").Set(reflect.ValueOf("second"))

// 	// fmt.Println(i)

// 	t := reflect.TypeOf(i)
// 	v := reflect.ValueOf(i)
// 	k := v.Kind()
// 	cnt := v.Len()

// 	fmt.Println(i)
// 	fmt.Println(t)
// 	fmt.Println(v)
// 	fmt.Println(k)
// 	fmt.Println(cnt)

// 	for i := 0; i < cnt; i++ {
// 		fmt.Println(v.Index(i))
// 	}

// 	// if t.Field(0).Name == "Name" && v.Field(0).CanSet() {
// 	// 	// v.Addr().Field(
// 	// 	// v.Field(0).SetString("second")
// 	// }

// 	// if v.Field(0).Kind() == reflect.String {
// 	// 	v.Field(0).SetString("second")
// 	// }
// 	// //v.Field(0).SetString("second")

// 	// fmt.Println(i)

// 	// if k == reflect.Struct {
// 	// 	cnt := t.NumField()
// 	// 	fmt.Println("count: ", cnt)

// 	// 	for i := 0; i < cnt; i++ {
// 	// 		//fmt.Println(i, " field is: ", v.Field(i), " type: ", v.Field(i).Type(), "name: ")
// 	// 		fmt.Println(i, ": ", t.Field(i).Name)
// 	// 	}
// 	// }

// 	// fmt.Println(t)
// 	// fmt.Println(v)
// 	// fmt.Println(k)
// 	// fmt.Println()
// }

// func main() {
// 	stu1 := Student{
// 		"first",
// 		1,
// 	}
// 	helper([]Student{stu1})
// }

// func main() {
// 	x := 20

// 	val := reflect.ValueOf(&x)

// 	fmt.Println(val.Elem())

// 	if val.Elem().CanSet() {
// 		val.Elem().Set(reflect.ValueOf(30))

// 		fmt.Println(reflect.ValueOf(x))
// 	}
// }
