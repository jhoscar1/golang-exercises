package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.side * s.side
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{
		radius: 1.5,
	}

	sq := square{
		side: 4,
	}

	fmt.Println(sq.area())
	fmt.Println(c1.area())
}
