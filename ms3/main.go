package main

import (
	"fmt"
	"os"
	"strconv"
)

// 注：Go 是“按值传递”编程语言。 每次向函数传递值时，Go 都会使用该值并创建本地副本（内存中的新变量）。
// 在函数中对该变量所做的更改都不会影响你向函数发送的更改。

func main() {
	sum, mul := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
	fmt.Println("Mul:", mul)

	firstName := "John"
	updateName(&firstName)
	fmt.Println(firstName)
}

func calc(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}

// 指针操作
func updateName(name *string) {
	*name = "David"
}
