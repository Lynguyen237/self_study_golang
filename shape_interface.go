package main

import (
	"fmt"
	"math"
)

type (
	geometry interface {
		area() float64
		perim() float64
	}
)

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) perim() float64 {
	return 2 * c.radius * math.Pi
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	shapes := []geometry{c, r}
	fmt.Println("==== Using shapes slice:")
	fmt.Println("Print data about shapes[0]: ", shapes[0])
	for _, shape := range shapes {
		fmt.Printf("%T: %v \n", shape, shape.area())
	}

	fmt.Println("==== Using measure function:")
	measure(r)
	measure(c)
}
