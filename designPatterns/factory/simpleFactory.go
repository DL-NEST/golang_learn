package main

import "fmt"

/* 简单工厂模式 */

const (
	square = iota
	cylindrical
)

// Model 一个工厂接口
type Model interface {
	enlarge()
}

// Cube 正方体
type Cube struct {
}

func (s Cube) enlarge() {
	fmt.Println("Cube enlarge")
}

// Cylindrical 圆柱体
type Cylindrical struct {
}

func (c Cylindrical) enlarge() {
	fmt.Println("Cylindrical enlarge")
}

func newModel(class uint) Model {
	switch class {
	case square:
		return Cube{}
	case cylindrical:
		return Cylindrical{}
	}
	return nil
}

func main() {
	model1 := newModel(square)
	model2 := newModel(cylindrical)
	model1.enlarge()
	model2.enlarge()
}
