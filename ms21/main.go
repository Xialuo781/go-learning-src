package main

import (
	"fmt"
	"time"
)

var quit = make(chan bool)

func fib(c chan int) {
	a, b := 1, 1

	for {
		// select 阻塞
		select {
		case c <- a:
			a, b = b, a+b
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func main() {
	start := time.Now()

	commond := ""
	data := make(chan int)

	go fib(data)

	for {
		num := <-data // 阻塞
		fmt.Println(num)
		fmt.Scanf("%s", &commond)
		if commond == "quit" {
			quit <- true
			break
		}
	}

	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
