package main

import "fmt"

/**
Switch 使用与其他语言不同的是：
1、左花括号（必须与switch 处于同一行）
2、条件表达式不限制为常量或者整数
3、单个case中，可以出现多个结果选项
4、不需要用break 来明确退出一个case
5、只有在case中明确添加 fallthrough 关键字，才会继续执行紧跟的下一个casse
6、 可以不设定switch 之后的条件表达式，在此种情况下，整个switch接口与多个if else 的逻辑作用等同
 */
func main() {
	switchTest2(5)
}

func switchTest(num int) {
	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fallthrough //只有在case中明确添加 fallthrough 关键字，才会继续执行紧跟的下一个casse
	case 4:
		fmt.Println("4")
	case 5,6,7:
		fmt.Println("5 6 7")
	}
}

func switchTest2(num int)  {
	switch  {
	case 0 <= num && num <= 4:
		fmt.Println("0 <= num && num <= 4")
	case 4 < num && num <= 10:
		fmt.Println("4 < num && num <= 10")
	}
}