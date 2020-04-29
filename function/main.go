package main

import "fmt"

/**
小写字母开头的函数只在本包内可见，大写字母开头的函数才能被其他包使用

*/
func main()  {
	var n int
	var num *int = &n

	Multiply(2,3,num)

	fmt.Println(*num)

	// 第一种测试 最大值
	maxNum := max(1,2,3,2,1,5,3,5,8,3,2,43,5,6)

	fmt.Println(maxNum)

	nums := []int{1,2,3,2,1,5,3,5,8,3,2,43,5,6}
	fmt.Printf("最大值 %d", max(nums...))
	Test()

}


func Multiply(a,b int, num *int) {
	*num = a * b
}

func max(nums ...int) int {
	max := nums[0]
	for _, val := range(nums) {
		if val > max {
			max = val
		}
	}
	return max
}