package main

import "fmt" // 格式化，输出/输入的方法

func main() {
	/*const (
		a = iota
		b
		c
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	*/
	// add()
	// getData()
	fuzhi()
}
func sushu() {
	var num int = 10 / 4 
	fmt.Println(num) 

	var num1 float64 = 10 / 4
	fmt.Println(num1) 

	var num2 float64 = 10.0 / 4
	fmt.Println(num2) 

	// %  => a % b = a - a/b*b 
	fmt.Println(10 % 3)  // 1
	fmt.Println(-10 % 3)  // -1
	fmt.Println(10 % -3)  // 1
	fmt.Println(-10 % -3)  //-1 

}

func fuzhi() {
	var num = 12 
	fmt.Println(num) // 12
	num1 := 1 +2 
	fmt.Println(num1)  // 3
	num += num1
	fmt.Println(num) // 15 num1= 3
	num -= num1
	fmt.Println(num) // num = 12 num1 = 3
	num *= num1
	fmt.Println(num) // num = 36 num1 = 3
	num /= num1
	fmt.Println(num) // num = 12 num1 = 3
	num %= num1
	fmt.Println(num)  //num = 0 num1 = 3

}




func add() {
	var num int = 12
	var num1 int // 默认0
	var num2 = 123.32
	num3 := 12.23

	var a, b, c = 12, 13, 14

	fmt.Println(num)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func getData() {
	var a *int
	num := 12
	a = &num // 取地址的时候必须要取一个变量的地址 赋值给指针
	fmt.Println(a)
}

func jis2() {
	// var a int
	//var b int32
	//a = 15
	//   b = a + a     // 编译错误
	//b = b + 5    // 因为 5 是常量，所以可以通过编译
	var n = 12
	fmt.Printf("n 的类型是 %T", n)
}
