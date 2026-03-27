package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

type Point struct {
	x, y float64
}

// NewPoint sets values of Point fields.
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// X gets x field of Point.
func (p Point) X() float64 {
	return p.x
}

// Y gets y field of Point.
func (p Point) Y() float64 {
	return p.y
}

// Distance calculates distance between point p and other.
func (p Point) Distance(other Point) float64 {
	return math.Sqrt((other.x-p.x)*(other.x-p.x) + (other.y-p.y)*(other.y-p.y))
}

func main() {
	x1 := float64(rand.IntN(100)) + rand.Float64()
	y1 := float64(rand.IntN(100)) + rand.Float64()
	x2 := float64(rand.IntN(100)) + rand.Float64()
	y2 := float64(rand.IntN(100)) + rand.Float64()

	p1 := NewPoint(x1, y1)
	p2 := NewPoint(x2, y2)
	fmt.Printf("Point1: (%f, %f), Point2: (%f, %f)\n",
		p1.X(), p1.Y(), p2.X(), p2.Y())

	res := p1.Distance(p2)
	fmt.Println("Result:", res)
}
