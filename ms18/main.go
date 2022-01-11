package main

import (
	"fmt"
)

// 可以指定channel的方向
func send(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}

func main() {
	size := 2
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	go send(ch, "three")
	go send(ch, "four")
	fmt.Println("All data send to the channel ...")

	for i := 0; i < 4; i++ {
		read(ch)
	}

	fmt.Println("Done!")
}
