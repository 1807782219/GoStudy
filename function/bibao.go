package main

import "fmt"
// 闭包的使用

func AddUpper() func (int) int {
	n := 0
	return func (x int) int {
		n = n + x
		return n 
	}
}

func main() {
	f := AddUpper()
	num := f(1)
	fmt.Println(num)
	
	f = AddUpper()
	num = f(1)
	fmt.Println(num)
	num = f(1)
	fmt.Println(num)
}