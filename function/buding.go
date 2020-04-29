package main

import "fmt"

func main()  {
	fun2(1,2,3,4)

	//测试全类型数据
	fun3(int(12), "123")
	/**
		fun3(int(12), "123") 结果如下：
		使用for - range 进行遍历：
		12 是一个int 数据
		123 是一个string 数据
	 */
	// 第二种测试传值： 发现传值错误
	slice := make([]interface{},2)
	slice[0] = 12
	slice[1] = 123
	fun3(slice)
	/**
	fun3(slice) 结果如下：
	使用for - range 进行遍历：
	[12 123] 是一个其他类型的数据
	发现直接传类型为[]interface{} Slice 的时候，只做为一个数据进行传值
	 */


}

func fun2(args... int)  {
	fmt.Println("不定参数的传值：全部数据传递")
	func1(args...) // 不定参数的传值： 按原样进行传递 // 1 	2 	3 	4
	fmt.Println("不定参数的传值：部分数据传递")
	func1(args[:3]...) // 1 	2 	3
}

/**
不定参数的函数
从内部实现来看， args... int 本质上一个数组切片
*/
func func1(args... int) {
	for _, arg := range args {
		fmt.Printf("%d \t", arg)
	}
	fmt.Println()
}

// interface{} 代表任意类型
func fun3(args... interface{})  {
	fmt.Println("使用for - range 进行遍历：")
	for _, arg := range args {
		switch arg.(type) { //arg.(type) 获取参数的类型
		case int:
			fmt.Println(arg, "是一个int 数据")
		case string:
			fmt.Println(arg, "是一个string 数据")
		default:
			fmt.Println(arg, "是一个其他类型的数据")
		}
	}
}

