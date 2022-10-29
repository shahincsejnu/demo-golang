package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shahincsejnu/demo-golang/class-8/math"
)

func main() {
	flg := flag.Int("temp", 100, "just a temp")

	flag.Parse()

	fmt.Println(*flg)

	Helper()

	math.Abs(100.2)

	gin.New()
}
