package main

import (
	"fmt"
	"main/src"
)

func main() {
	c := src.Compute{
		D: 10,
	}
	res := c.AddTwo(1, 2)
	fmt.Println(res)

	res2, err := c.Divide(10, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(res2)

	res3, err := c.Divide(10, -2)
	if err != nil {
		panic(err)
	}
	fmt.Println(res3)
}
