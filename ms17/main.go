package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	// // 使用 make() 函数创建 channel 时，会创建一个无缓冲 channel，这是默认行为
	// // 无缓冲 channel 会阻塞发送操作，直到有人准备好接收数据。 这就是为什么我们之前说发送和接收都属于阻塞操作
	// ch := make(chan string)

	// 有缓冲的channel
	ch := make(chan string, 5)

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	// 无缓冲 channel 的发送和接收操作是同步的。 即使使用并发，通信也是同步的
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
