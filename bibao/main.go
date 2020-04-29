package main

import "fmt"

/**
匿名函数和闭包函数
 */
func main() {
	sum := nimingTest(1,2 )
	fmt.Println(sum)

	// 闭包函数实践
	func1 := addUper()
	fmt.Println(func1(1)) //1
	fmt.Println(func1(2)) //3
	fmt.Println(func1(3)) //6

	func2 := addUper()
	fmt.Println(func2(1)) //1
	fmt.Println(func2(2)) //3
	fmt.Println(func2(3)) //6
}


/**
要说闭包，首先需要说匿名函数。匿名函数顾名思义就是把函数名称给隐藏了、不要了，只有函数的关键字，参数列表、返回值列表和实现
 */
func nimingTest(a, b int) int {
	sum := func(a ,b int) int {
		return a + b;
	}
	return sum(a,b)
}

/**
下面是闭包：闭包就是一个函数和其相关的引用环境组合的一个整体
简单来说就是 一个匿名函数和这个匿名函数中使用到的外部参数形成了一个整体。比如：闭包是类， 匿名函数是类中的方法，而 外部参数就是 类中的变量
 */

func addUper() func(x int) int {
	n := 0
	return func(x int) int {
		n = n + x
		return n
	}
}



