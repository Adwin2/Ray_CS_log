package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) //function call
	fmt.Println(p.Distance(q))  // method call

}

// traditional func
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing but as a method of Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
