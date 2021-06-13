package main

import "fmt"

type shape interface {
	printArea()
}

type square struct {
	sideLength float64
}

type traingle struct {
	base   float64
	height float64
}

func (s square) printArea() {
	fmt.Println("The area of sqaure is", s.sideLength*s.sideLength)
}

func (t traingle) printArea() {
	fmt.Println("The area of triangle is", t.height*t.base*0.5)
}

func printArea(s shape) {
	s.printArea()
}
func main() {
	tr := traingle{
		3,
		4,
	}
	sq := square{
		sideLength: 4,
	}
	printArea(sq)
	printArea(tr)
}
