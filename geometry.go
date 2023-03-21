package geometry

import (
	"fmt"
	"math"
)

const pi = math.Pi

type Geometry interface {
	Area() float64
	Perimetr() float64
	GetArgs() string
}

type Rectangle struct {
	a float64
	b float64
}

func NewRectangle(a float64, b float64) Rectangle {
	return Rectangle{a, b}
}

func (r Rectangle) Area() float64 {
	return r.a * r.b
}
func (r Rectangle) GetArgs() string {
	s := fmt.Sprintf("стороны %v и %v", r.a, r.b)
	return s
}
func (r Rectangle) Perimetr() float64 {
	return 2*r.a + 2*r.b
}

type Triangle struct {
	a float64
	b float64
	c float64
	h float64
}

func NewTriangle(a float64, b float64, c float64, h float64) Triangle {
	return Triangle{a, b, c, h}
}

func (t Triangle) Area() float64 {
	return t.a * t.h
}

func (t Triangle) Perimetr() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) GetArgs() string {
	s := fmt.Sprintf("стороны %v, %v и %v; высота - %v; ", t.a, t.b, t.c, t.h)
	return s
}

type Circle struct {
	r float64
}

func NewCircle(r float64) Circle {
	return Circle{r}
}
func (c Circle) GetArgs() string {
	s := fmt.Sprintf("радиус - %v", c.r)
	return s
}

func (c Circle) Area() float64 {
	return pi * c.r * c.r
}

func (c Circle) Perimetr() float64 {
	return 2 * pi * c.r
}
