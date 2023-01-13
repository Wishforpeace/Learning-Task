package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
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

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	p = Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Println(p)
	p.ScaleBy(2)
	fmt.Println(p)

	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.X)
	cp.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}

	blue := color.RGBA{0, 0, 255, 255}

	var p1 = ColoredPoint{Point{1, 1}, red}

	var q1 = ColoredPoint{Point{5, 4}, blue}

	fmt.Println(p1.Distance(q1.Point))

	p1.ScaleBy(2)

	q1.ScaleBy(2)

	fmt.Println(p1.Distance(q1.Point))
}
