package main

import (
	"fmt"

	"github.com/mail2sada/training/math"
)

func main() {
	fmt.Println("Demo: package publish")
	res := math.Add(10, 20)
	fmt.Println(res)

	res = math.Sub(200, 100)

	fmt.Println(res)

}
