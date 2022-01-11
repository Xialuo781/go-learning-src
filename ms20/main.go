package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float64, ch chan float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- x
}

func main() {
	start := time.Now()

	size := 15
	ch := make(chan float64, size)

	for i := 1; i < size; i++ {
		go fib(float64(i), ch)
	}

	for i := 1; i < size; i++ {
		n := <-ch
		fmt.Printf("Fib(%v): %v\n", i, n)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
