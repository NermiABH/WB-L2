package main

import (
	"fmt"
	"math"
)

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForTriangle(*Triangle)
}

type Square struct {
	Side float64
}

func (s *Square) accept(visitor Visitor) {
	visitor.visitForSquare(s)
}

type Circle struct {
	Radius float64
}

func (c *Circle) accept(visitor Visitor) {
	visitor.visitForCircle(c)
}

type Triangle struct {
	SideA, SideB, SideC float64
}

func (t *Triangle) accept(visitor Visitor) {
	visitor.visitForTriangle(t)

}

type PerimeterCalculate struct {
	perimeter float64
}

func (p *PerimeterCalculate) visitForSquare(square *Square) {
	p.perimeter = square.Side * 4
}

func (p *PerimeterCalculate) visitForCircle(circle *Circle) {
	p.perimeter = 2 * math.Pi * circle.Radius
}

func (p *PerimeterCalculate) visitForTriangle(triangle *Triangle) {
	p.perimeter = triangle.SideA + triangle.SideB + triangle.SideC
}

type AreaCalculate struct {
	area float64
}

func (s *AreaCalculate) visitForSquare(square *Square) {
	s.area = square.Side * square.Side
}

func (s *AreaCalculate) visitForCircle(circle *Circle) {
	s.area = math.Pi * circle.Radius * circle.Radius
}

func (s *AreaCalculate) visitForTriangle(triangle *Triangle) {
	p := (triangle.SideA + triangle.SideB + triangle.SideC) / 2
	s.area = math.Sqrt(p * (p - triangle.SideA) * (p - triangle.SideB) * (p - triangle.SideC))
}

func main() {
	square := &Square{Side: 10}
	circle := &Circle{Radius: 5}
	triangle := &Triangle{SideA: 3, SideB: 4, SideC: 5}

	perimeterCalculate := &PerimeterCalculate{}
	square.accept(perimeterCalculate)
	fmt.Printf("Square: Perimeter %v\n", perimeterCalculate.perimeter)

	circle.accept(perimeterCalculate)
	fmt.Printf("Circle: Perimeter %v\n", perimeterCalculate.perimeter)

	triangle.accept(perimeterCalculate)
	fmt.Printf("Triangle: Perimeter %v\n", perimeterCalculate.perimeter)

	fmt.Println()
	areaCalculate := &AreaCalculate{}
	square.accept(areaCalculate)
	fmt.Printf("Square: Area %v\n", areaCalculate.area)

	circle.accept(areaCalculate)
	fmt.Printf("Circle: Area %v\n", areaCalculate.area)

	triangle.accept(areaCalculate)
	fmt.Printf("Triangle: Area %v\n", areaCalculate.area)
}
