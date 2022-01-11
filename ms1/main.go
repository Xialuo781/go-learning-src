package main

import "fmt"

var OutFuncValue0 string

var (
	OutFuncValue1, OutFuncValue2 = "OutFuncValue1", "OutFuncValue2"
)

const HTTPStatusOK = 200

const (
	StatusOK              = 0
	StatusConnectionReset = 1
	StatusOtherError      = 2
)

func main() {
	firstName, LastName := "xia", "luo"
	age := 32
	fmt.Println(firstName, LastName, age)
}
