package main

import "fmt"

type Demo interface {
	Sum(int, int) (int, error)
}

type Temp struct {
}

func (t Temp) Sum(x, y int) (int, error) {
	return x + y, nil
}

func process(d Demo, x, y int) (int, error) {
	res, err := d.Sum(x, y)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func main() {
	temp1 := Temp{}

	res, err := process(temp1, 1, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
