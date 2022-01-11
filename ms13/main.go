package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.SetPrefix("main(): ")
	log.Print("Hey, I'm a log!")
	log.Fatal("Hey, I'm an error log!")
	// log.Panic("Hey, I'm an error log!")
	fmt.Print("Can you see me?")
}
