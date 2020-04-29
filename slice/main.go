package main

import "fmt"

/**
切片的操作
 */
func main() {
	func4()
}


/**
数组切片可以基于一个已存在的数组创建，数组切片可以值使用数组的一部分元素或者整个数组来创建，甚至可以创建一个比所基于的数组还要打的数组切片
 */
func func1() {
	array := []int{1,2,3,4,5,6,7}
	fmt.Println("数组的值：")
	for _,v := range array {
		fmt.Printf("%d \t",v)
	}
	fmt.Println("\n 切片1的值：")
	slice1 := array[1:4]    //index = start ~ end-1 => 2,3,
	for _,v := range slice1 {
		fmt.Printf("%d \t",v)
	}

	fmt.Println("\n 切片2的值：")
	slice2 := array[:]    //index = start ~ end-1 => 2,3,
	for _,v := range slice2 {
		fmt.Printf("%d \t",v)
	}
}

/**
直接创建：
Go 语言提供内置的函数make() 可以用于创建数组切片，
slice := make([]int, 5)
创建好的切片初始值都是 0， 共计 5个
 */
func func2() {
	slice := make([]string, 5)
	slice1 := make([]string, 5, 10)
	slice2 := []string{"1","2","3"}
	for _, v := range slice {
		fmt.Println(v)
	}
	fmt.Printf("切片的容量：%d， 切片的长度为 %d\n", cap(slice), len(slice))
	fmt.Printf("切片1的容量：%d， 切片的长度为 %d\n", cap(slice1), len(slice1))
	fmt.Printf("切片2的容量：%d， 切片的长度为 %d\n", cap(slice2), len(slice2))

	for i:= 1 ;i< len(slice) - 2;i++ {
		slice[i] = "A" + string(99+i)
	}

	fmt.Printf("切片的容量：%d， 切片的长度为 %d\n", cap(slice), len(slice))

	for _, v := range slice {
		fmt.Println(v)
	}
}

/**
切片的动态新增数据操作
 */
func func3() {
	silce := make([]int, 5, 10)
	fmt.Println(silce)
	// 第一次新增
	silce = append(silce, 1,2,3)
	fmt.Println(silce)
	fmt.Printf("第一次新增后，切片的容量：%d， 切片的长度为 %d\n", cap(silce), len(silce))// 第一次新增后，切片的容量：10， 切片的长度为 8
	// 第二次新增 : 直接增加一个切片中的内容
	slice2 := []int{4,5}
	silce  = append(silce, slice2...)
	fmt.Println(silce)
	fmt.Printf("第二次新增后，切片的容量：%d， 切片的长度为 %d\n", cap(silce), len(silce))//第二次新增后，切片的容量：10， 切片的长度为 10
	// 第三次新增时，需要切片进行扩容才能存储
	slice3 := []int{6,7,8,9,10}
	silce  = append(silce, slice3...)
	fmt.Println(silce)
	fmt.Printf("第三次新增后，切片的容量：%d， 切片的长度为 %d\n", cap(silce), len(silce))//第三次新增后，切片的容量：20， 切片的长度为 15
}

func func4() {
	slice_from :=[]int{1,2,3,4} //[1 2 3 4]
	slice_to := make([]int, 10) //[0 0 0 0 0 0 0 0 0 0]

	n:= copy(slice_to, slice_from)
	fmt.Println(slice_to) //[1 2 3 4 0 0 0 0 0 0]
	fmt.Printf("Copied %d elements\n", n) //Copied 4 elements


	n = copy(slice_from, slice_to)
	fmt.Println(slice_from) //[1 2 3 4]
	fmt.Printf("Copied %d elements\n", n) //Copied 4 elements
}