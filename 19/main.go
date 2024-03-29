package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s *square) area() float64 {
	return s.length * s.length
}

func (c *circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	circleArea := circle{
		radius: 20,
	}

	squareArea := square{
		length: 10,
	}

	info(&circleArea)
	info(&squareArea)
}
