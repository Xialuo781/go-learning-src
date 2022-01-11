package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

// type Employee0 struct {
// 	Information Person
// 	ManagerID   int
// }

type Employee struct {
	Person
	ManagerID int
}

func main() {
	// person := Person{1001, "John", "Doe"}
	// // person2 := Person{LastName: "Doe", FirstName: "John"}
	// person.ID = 1002
	// fmt.Println(person)
	// fmt.Println(person.FirstName)

	// // 使用取地址复制结构
	// personCopy := &person
	// personCopy.FirstName = "David"
	// fmt.Println(person.FirstName)

	// var employee Employee0
	// employee.Information.FirstName = "John"
	// 嵌套直接展开
	employee := Employee{Person: Person{FirstName: "John"}}
	employee.LastName = "Doe"

	// 对Json的支持
	employees := [...]Employee{
		Employee{Person: Person{LastName: "Doe", FirstName: "John"}},
		Employee{Person: Person{LastName: "Campbell", FirstName: "David"}},
	}
	data, _ := json.Marshal(employees)
	fmt.Printf("%s \n", data)

	var decode []Employee
	json.Unmarshal(data, &decode)
	fmt.Printf("%v", decode)
}
