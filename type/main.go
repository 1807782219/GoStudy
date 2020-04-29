package main

import "fmt"

type Integer int

// 获取两个数的最小值
func (a Integer) Less(b Integer) Integer {
	if a > b {
		return b
	}
	return a
}

func (a Integer) Add(b Integer) {
	a += b
}

func (a *Integer) AddWith(b Integer) {
	*a += b
}

func main() {
	var a Integer = 2
	//获取两个数的最小值
	fmt.Println(a.Less(1))

	// 新增方式1
	a.Add(1)
	fmt.Println(a) //2

	// 新增方式2
	a.AddWith(1)
	fmt.Println(a) //3
	/**
	GO 语言中没有隐藏的this 指针，含义可理解为：
	1. 方法施加的目标（也就是“对象”） 显示传递，没有被隐藏起来
	2. 方法施加的目标（也就是“对象”）不需要非的是指针，也不用非得叫this

	GO 语言和C语言都是基于值传递的，如果想修改变量的值，只能传递指针
	 */
}