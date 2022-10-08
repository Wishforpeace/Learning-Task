package main

import "fmt"

type Computer interface {
	compute(int, int) interface{}
}

type UseComputer struct{}
type Add struct{}
type Subtract struct{}
type Multiply struct{}
type Divide struct{}

func (Add *Add) compute(a int, b int) interface{} {
	return a + b
}
func (Sub *Subtract) compute(a int, b int) interface{} {
	return a - b
}
func (Mul *Multiply) compute(a int, b int) interface{} {
	return a * b
}
func (Div *Divide) compute(a int, b int) interface{} {
	return float64(a) / float64(b)
}

func (u *UseComputer) useCom(com Computer, one int, two int) interface{} {
	return com.compute(one, two)
}

func main() {
	var (
		add Add
		sub Subtract
		mul Multiply
		div Divide
	)

	var UC UseComputer
	one := 13
	two := 2

	fmt.Println("加法 ", UC.useCom(&add, one, two))

	fmt.Println("减法 ", UC.useCom(&sub, one, two))

	fmt.Println("乘法 ", UC.useCom(&mul, one, two))

	fmt.Println("除法 ", UC.useCom(&div, one, two))
}
