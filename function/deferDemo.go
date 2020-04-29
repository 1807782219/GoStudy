package main

import (
	"fmt"
	f "goStudy/utils" // 包名一般是文件夹名称，一个文件夹可以有多个文件，且这些文件同属于一个包
)
func main()  {
	
	// deferDemo()
	sum := f.Callback(2, f.Add)
	println(sum) // 5

	f.TestOther()

}

func deferDemo() {
	fmt.Println("111111111111")
	defer fmt.Println("3333")
	n := 444444
	defer fmt.Println(n)
	fmt.Println("222222222222")
	/*
		结果：
		111111111111
		222222222222
		444444
		3333
		： 当多个defer 的时候，多个defer 通过栈 来存储
	*/

}

