package main

import (
	"fmt"
	"math"
)

type Pointers struct {
	X int
	Y int
}

type PointersBuilderI interface {
	SetX(value int) PointersBuilderI
	SetY(value int) PointersBuilderI

	Build() Pointers
}

type pointersBuilder struct {
	x int
	y int
}

func (p *pointersBuilder) SetX(value int) PointersBuilderI {
	p.x = value

	return p
}

func (p *pointersBuilder) SetY(value int) PointersBuilderI {
	p.y = value

	return p
}

func (p pointersBuilder) Build() Pointers {
	return Pointers{
		X: p.x,
		Y: p.y,
	}
}

func NewPointersBuilder() pointersBuilder {
	return pointersBuilder{}
}

func CalculateDistance(p1 pointersBuilder, p2 pointersBuilder) float64 {
	result := math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2))
	return result
}

func main() {
	pointer1 := NewPointersBuilder()
	pointer1.SetX(1).SetY(1).Build()
	pointer2 := NewPointersBuilder()
	pointer2.SetX(2).SetY(2).Build()

	fmt.Println(CalculateDistance(pointer1, pointer2))
}
