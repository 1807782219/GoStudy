package main

import "fmt"

/**
init() 函数，特点： 在main函数之前调用
			使用场景：初始化数据库环境等
defer 关键字，特点： defer关键字修饰的内容在main函数结束之前调用，或者出现panic等错误的时候在函数结束之前调用。
					采用栈的特性，先进后出，逐一执行
			 使用场景：打开的文件流关闭操作，出现异常时的处理
 */
func init()  {
	fmt.Println("init 函数。。。。")
}

func main() {
	defer fmt.Println("defer 1 ...")
	fmt.Println("main .....")
	defer fmt.Println("defer 2 ...")

	/**
	结果：
	init 函数。。。。
	main .....
	defer 2 ...
	defer 1 ...
	 */

}
