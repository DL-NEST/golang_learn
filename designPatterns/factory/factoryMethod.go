package main

import "fmt"

/* 工厂方法模式 */

const (
	circular = iota
	triangle
)

// ShapeFactory 工厂的接口
type ShapeFactory interface {
	CreateShape() Shape
}

type Shape interface {
	Enlarge(data string)
}

type Circular struct {
}

func (j Circular) Enlarge(data string) {
	fmt.Printf("Circular %s\n", data)
}

type Triangle struct {
}

func (y Triangle) Enlarge(data string) {
	fmt.Printf("Triangle %s\n", data)
}

type CircularFactory struct {
}

func (j CircularFactory) CreateShape() Shape {
	// 这里可以对Shape进行组装
	return Circular{}
}

type TriangleFactory struct {
}

func (y TriangleFactory) CreateShape() Shape {
	//这里可以对parser做特殊处理
	return Triangle{}
}

// NewShapeFactory 对象创建
func NewShapeFactory(shape int) ShapeFactory {
	switch shape {
	case circular:
		return CircularFactory{}
	case triangle:
		return TriangleFactory{}
	}
	return nil
}

func main() {
	circular := NewShapeFactory(circular).CreateShape()
	triangle := NewShapeFactory(triangle).CreateShape()
	circular.Enlarge("创建")
	triangle.Enlarge("创建")
}
