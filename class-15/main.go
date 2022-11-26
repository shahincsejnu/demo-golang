package main

import (
	"fmt"
)

type FileSystem interface {
	Read()
	Write()
	List()
}

type Linode struct {
	config *config
}

func (a AWS) Read() {

}

func (a AWS) Write() {

}

func (a AWS) List() {

}

type AWS struct {
	config *config
}

func (l Linode) Read() {

}

func (l Linode) Write() {

}

func (l Linode) List() {

}

type Student struct {
	Name string
	Id   int
}

func (s Student) demo3(x int) {
	fmt.Println(s.Id)
	fmt.Println(s.Name)
}

func main() {
	var fs []FileSystem = []FileSystem{Linode{}, AWS{}}

	for _, f := range fs {
		f.List()
	}

	stu1 := Student{
		Name: "jahid",
		Id:   1,
	}

	stu1.demo3(10)

	// ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(1*time.Millisecond))
	// defer cancel()

	// req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://google.com", nil)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// defer res.Body.Close()

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("size :", len(data))

	// demo()
}

func demo() {

}

func demo2() {
	demo()
}
