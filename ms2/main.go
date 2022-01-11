package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整型
	// var integer8 int8 = 127
	// var integer16 int16 = 32767
	// var interger32 int32 = 2147483647
	// var integer64 int64 = 9223372036854775807
	// fmt.Println(integer8, integer16, interger32, integer64)

	// rune
	// rune := 'G'
	// fmt.Println(rune)

	// 浮点型
	// var float32 float32 = 2147483647
	// var float64 float64 = 9223372036854775807
	// fmt.Println(float32, float64)
	// fmt.Println(math.MaxFloat32, math.MaxFloat64)

	// 字符串
	// var firstName string = "John"
	// lastName := "Doe"
	// fmt.Println(firstName, lastName)

	// fullName := "John Doe \t(alias \"Foo\")\n"
	// fmt.Println(fullName)

	// 强制类型转换
	var integer16 int16 = 127
	var integer32 int32 = 32767
	fmt.Println(int32(integer16) + integer32)

	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
