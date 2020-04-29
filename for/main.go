package main

import "fmt"


func main() {
	// jizita(9)
	func2()
}


func func2() {
	var nums [5]int64

	for i:= 0;i<len(nums);i++ {
		fmt.Printf("请输入第%v 的值：", i+1)
		fmt.Scanln(&nums[i])
	}

	fmt.Println("开始打印。。。。")
	for _,v := range nums {
		fmt.Println(v)
	}

	var nums2 = [...]int64{1,2,3,4}

	for index, value := range nums2 {
		fmt.Printf("第%d的元素的值为%v\n", index + 1, value)
	}
}

// 打出空心三角
func jizita(m  int) {
	for i := 1; i <= m; i++ {
		for k := 1; k <= m - i; k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2 * i - 1; j++ {
			if j == 1 || j == 2*i -1 || i == m{
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
			
		}
		fmt.Println("")
	}

}
// 打印实心三角
func jizita1(m  int) {
	for i := 1; i <= m; i++ {
		for k := 1; k <= m - i; k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2 * i - 1; j++ {	
			fmt.Printf("*")	
		}
		fmt.Println("")
	}

}
func forDemo() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	num := 10
	for num >= 1 {
		fmt.Println(num)
		num = num - 1
	}

	str := "hello world"

	for index, value := range str {
		fmt.Printf("Index :%d , Value is %c \n", index, value)
	}
}
