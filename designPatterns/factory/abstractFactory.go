package main

import "fmt"

/* 抽象工厂模式 */

// Model1 一个工厂接口
type Model1 interface {
	enlarge()
}

// Cube1 正方体
type Cube1 struct {
}

func (s Cube1) enlarge() {
	fmt.Println("Cube enlarge")
}

// Cylindrical1 圆柱体
type Cylindrical1 struct {
}

func (c Cylindrical1) enlarge() {
	fmt.Println("Cylindrical enlarge")
}

type IModelFactory interface {
	Cube() Cube1
	Cylindrical() Cylindrical1
}

type ModelFactory struct {
}

func (m ModelFactory) Cube() Cube1 {
	return Cube1{}
}

func (m ModelFactory) Cylindrical() Cylindrical1 {
	return Cylindrical1{}
}

func main() {
	modelFactory := ModelFactory{}
	modelFactory.Cylindrical().enlarge()
	modelFactory.Cube().enlarge()
}
