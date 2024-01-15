package src

import "fmt"

type Compute struct {
	D int // d is a private field, D is a public field
}

type ICompute interface {
	AddTwo(a int, b int) int
	Divide(a int, b int) (int, error)
}

func (c *Compute) AddTwo(a int, b int) int {
	return a + b + c.D
}

func (c *Compute) Divide(a int, b int) (int, error) {
	if b <= 0 {
		return 0, fmt.Errorf("b must be greater than 0")
	}
	return a / b, nil
}
