package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// testDefer()

	// testPanic()

	testRecover()
}

func testDefer() {
	// defer 后进先出
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}

	// 打开文件后提前关闭
	newfile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}

	defer newfile.Close()

	if _, error = io.WriteString(newfile, "Learning Go!"); error != nil {
		fmt.Println("Error: Could not write to file")
		return
	}

	newfile.Sync()
}

func testPanic() {
	// panic
	highlow(2, 0)
	fmt.Println("Program finished successfully!") // 在highlow中发生了panic，此处不会打印
}

func testRecover() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")
}

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		// 此处，控制流中断，所有推迟的函数都开始输出“Deferred...”消息。
		panic("highlow() low greater than high") // 程序崩溃，并显示完整的堆栈跟踪。
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}
