package main

import "fmt"

var a, b, c int

func Test() (int, int) {
	var d, e, j int
	fmt.Println(j)
	return d, e
}
func main() {
	var f *int
	f = &a
	var h *int
	h = &b

	d, e := Test()

	var dd *int
	dd = &d

	var ee *int
	ee = &e

}
