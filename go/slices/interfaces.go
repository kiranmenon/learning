package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	h ,l float64
}

type circle struct {
	r float64
}

func (r rect) area() float64 {
	return r.h*r.l
}

func (r rect) perim() float64 {
	return 2*r.h + 2*r.l
}

func (c circle) area() float64  {
	return math.Pi * c.r*c.r
}

//func (c circle) perim() float64  {
//	return 2 * math.Pi * c.r
//}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	//fmt.Println(g.perim())
}

func main() {
	r := rect{12,8}
	c := circle{r: 10}
	measure(r)
	measure(c)
}
