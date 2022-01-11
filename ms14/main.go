package main

import (
	"fmt"
	"strings"
)

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

func (s square) perimeter() int {
	return s.size * 4
}

// 基于基本类型创建自定义类型，然后将其用作基本类型
type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

// 嵌入方法
// coloredTriangle对象同样具有perimeter函数，Go 编译器会通过创建包装器方法来推广 perimeter() 方法
type coloredTriangle struct {
	triangle
	color string
}

// 重载方法
func (ct coloredTriangle) perimeter() int {
	return ct.size * 3 * 2
}

func main() {
	// t := triangle{3}
	// t.doubleSize()
	// fmt.Println("Size:", t.size)
	// s := square{size: 3}

	// fmt.Println("Perimeter (triangle):", t.perimeter())
	// fmt.Println("Perimeter (square):", s.perimeter())

	s := upperstring("Learning Go!")
	fmt.Println(s)
	fmt.Println(s.Upper())

	ct := coloredTriangle{triangle: triangle{3}, color: "blue"}
	fmt.Println("Size:", ct.size)
	// 使用重载的函数
	fmt.Println("Perimeter (colored)", ct.perimeter())
	// 显示访问重载前的函数
	fmt.Println("Perimeter (normal)", ct.triangle.perimeter())
}
