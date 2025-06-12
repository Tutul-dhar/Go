package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float32
}

type rectangle struct {
	width, height float32
}

type circle struct {
	radius float32
}

func (r rectangle) Area() float32 {
	return r.height*r.width
} 

func (c circle) Area() float32 {
	return math.Pi * c.radius*c.radius
}

func calculateArea(s shape) float32 {
	return s.Area()
}

func main() {
	rect := rectangle{5.0, 4.0}
	cir := circle{4}

	shapes := []shape{&rect,&cir}

	for _, shape := range shapes {
		fmt.Println(calculateArea(shape))
	}
}