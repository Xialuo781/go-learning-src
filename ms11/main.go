package main

import (
	"fmt"
)

func fibonacci(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}
	res := make([]int, n)
	res[0], res[1] = 1, 1
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res
}

func romanToArabic(numeral string) int {
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	arabicVals := make([]int, len(numeral)+1)

	for index, digit := range numeral {
		if val, present := romanMap[digit]; present {
			arabicVals[index] = val
		} else {
			fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", numeral, digit)
			return 0
		}
	}

	total := 0

	for index := 0; index < len(numeral); index++ {
		if arabicVals[index] < arabicVals[index+1] {
			arabicVals[index] = -arabicVals[index]
		}
		total += arabicVals[index]
	}
	return total
}

func main() {
	// res := fibonacci(6)
	// fmt.Println(res)

	res := romanToArabic("MCM")
	fmt.Println(res)
}
