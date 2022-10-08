package main

import (
	"fmt"
	"math"
)

type Printer interface {
	PrintMyWay()
}
type Graph struct{}

func (graph *Graph) Print(printer Printer) {
	printer.PrintMyWay()
}

type Rectangle struct {
	Length    int
	Width     int
	Area      int
	Perimeter int
}
type Square struct {
	Rectangle
	Diagonal float64
}

func (rec *Rectangle) PrintMyWay() {
	rec.CalculateArea()
	rec.CalculatePerimeter()
	fmt.Printf("长:%d  宽:%d\n", rec.Length, rec.Width)
	fmt.Printf("面积:%d\n", rec.Area)
	fmt.Printf("周长:%d\n", rec.Perimeter)
}

func (sq *Square) PrintMyWay() {
	sq.CalculateArea()
	sq.CalculatePerimeter()
	sq.CalculateDiagonal()
	fmt.Printf("长:%d  宽:%d\n", sq.Length, sq.Width)
	fmt.Printf("面积:%d\n", sq.Area)
	fmt.Printf("周长:%d\n", sq.Perimeter)
	fmt.Printf("对角线:%f\n", sq.Diagonal)
}
func (rec *Rectangle) CalculateArea() {
	rec.Area = rec.Length * rec.Width
}
func (rec *Rectangle) CalculatePerimeter() {
	rec.Perimeter = 2 * (rec.Length + rec.Width)
}

func (sq *Square) CalculateDiagonal() {
	pow := float64(sq.Length*sq.Length + sq.Width*sq.Width)
	sq.Diagonal = math.Sqrt(pow)
}

func main() {
	var graph Graph
	var rec = Rectangle{}
	fmt.Println("请输入长 ")
	fmt.Scanf("%d", &rec.Length)
	fmt.Println("宽 ")
	fmt.Scanf("%d", &rec.Width)
	if rec.Length == rec.Width {
		var sq Square
		sq.Rectangle = rec
		graph.Print(&sq)
		return
	}
	graph.Print(&rec)
}
