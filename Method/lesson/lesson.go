package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}
type Point2 struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum

}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())

	fmt.Println(p)
	p.ScaleBy(2)
	fmt.Println(p)

	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.X)
	cp.Y = 2
	fmt.Println(cp.Y)

	cp2 := ColoredPoint{
		Point: Point{
			X: 2,
			Y: 4,
		},
	}
	fmt.Println(cp.Distance(cp2.Point))
}
