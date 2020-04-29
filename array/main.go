package main

import "fmt"

/*
数组的长度一旦定义后，不能再修改
数组是值类型，每次传递的时候传的都是数组的一个副本
 */

func main() {
	arrayTest3()
}

// 第一种赋值：初始化
func arrayTest1() {
	array := []int64{1,2,3,4,5}

	for i:= 0;i<len(array);i++ {
		fmt.Printf("%d \t", array[i])
	}
}
// 第二种赋值：每个赋值
func arrayTest2() {
	var array [5]int64

	for i:= 0;i<len(array);i++ {
		array[i] = int64(i + 1)
	}

	for _, arr := range array {
		fmt.Printf("%d \t", arr)
	}
}

// 第三种赋值：一起赋值
func arrayTest3() {
	var array []int64
	array2 := []int64{1,2,3,4}
	array = array2
	for _, arr := range array {
		fmt.Printf("%d \t", arr)
	}
}
