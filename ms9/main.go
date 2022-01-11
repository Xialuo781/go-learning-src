package main

import "fmt"

func main() {
	// 数组 //
	// var a [3]int
	// a[1] = 10
	// fmt.Println(a[0])
	// fmt.Println(a[1])
	// fmt.Println(a[len(a)-1])

	// cities := [5]string{"haha", "xixi", "hehe", "lala", "ouou"}
	// fmt.Println(cities)

	// q := [...]int{1, 2, 3}
	// fmt.Println(q)

	// // 仅指定最后一位
	// numbers := [...]int{99: -1}
	// fmt.Println("First Position:", numbers[0])
	// fmt.Println("Last Position:", numbers[99])
	// fmt.Println("Length: ", len(numbers))

	// // 多维数组
	// var twoD [3][5]int
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		twoD[i][j] = (i + 1) * (j + 1)
	// 	}
	// 	fmt.Println("Row", i, twoD[i])
	// }
	// fmt.Println("\nAll at once:", twoD)

	// 切片 //
	// months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	// fmt.Println(months)
	// fmt.Println("Length:", len(months))
	// fmt.Println("Capacity:", cap(months))
	// quarter1 := months[0:3]
	// quarter2 := months[3:6]
	// quarter3 := months[6:9]
	// quarter4 := months[9:12]
	// fmt.Println(quarter1, len(quarter1), cap(quarter1))
	// fmt.Println(quarter2, len(quarter2), cap(quarter2))
	// fmt.Println(quarter3, len(quarter3), cap(quarter3))
	// fmt.Println(quarter4, len(quarter4), cap(quarter4))

	// // 切片扩展
	// quarter2Extended := quarter2[:4]
	// fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))

	// // 追加项
	// var numbers []int
	// for i := 0; i < 10; i++ {
	// 	numbers = append(numbers, i)
	// 	fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	// }

	// // 删除项
	// letters := []string{"A", "B", "C", "D", "E"}
	// remove := 2

	// if remove < len(letters) {
	// 	fmt.Println("Before", letters, "Remove", letters[remove])
	// 	letters = append(letters[:remove], letters[remove+1:]...)
	// 	fmt.Println("After", letters)
	// }

	// // 用make创建切片，作为切片的副本
	// slice2 := make([]string, 3)
	// copy(slice2, letters[1:4])

	// 映射 //
	// studentsAge := map[string]int{
	// 	"john": 32,
	// 	"bob":  31,
	// }
	// fmt.Println(studentsAge)

	// 用make新建映射
	// var studentsAge map[string]int // 这样不行
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println(studentsAge)
	// 访问不存在的项会返回默认值
	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}

	// 删除项
	// delete(studentsAge, "john")
	fmt.Println(studentsAge)
	// 删除不存在的项不会报出panic

	// 遍历映射
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
	// for name := range studentsAge {
	//     fmt.Printf("Names %s\n", name)
	// }
}
