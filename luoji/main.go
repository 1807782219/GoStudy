
package main

import "fmt"

func main() {
	// forDemo()
	switchDemo()
}

func switchDemo() {
	var age int32
	fmt.Println("switch 1")
	fmt.Println("请输入一个数字")
	fmt.Scanln(&age)
	switch  {
		case 0<age && age < 18: 
			fmt.Println("未成年人")
		
		case 18 <= age && age < 30: 
			fmt.Println("青年人")
		case 30 <= age && age < 60:
			fmt.Println("中年人")
		case age >= 60:
			fmt.Println("老年人")
		default:
			fmt.Println("数据输入错误！")
	}
	fmt.Println("switch 2")
	switch a := 1 + 2; { //; 必须要加
	case a < 0:
		fmt.Println("小于0")
	case a == 0:
		fmt.Println("等于 0 ")
	default: 
		fmt.Println("大于0")
	}
	fmt.Println("switch 3")
	switch a := usedSwitchDemo(); { //; 必须要加
	case a < 0:
		fmt.Println("小于0")
	case a == 0:
		fmt.Println("等于 0 ")
	default: 
		fmt.Println("大于0")
	}

	fmt.Println("switch 4")
	switch age {
	case 20,40,50:
		fmt.Println("好时光")
	default:
		fmt.Println("HHHHHH")
	}


}

func usedSwitchDemo() int {
	return 1 + 2
}

func forDemo() {
	// for 1
	fmt.Println("for 1")
	for i := 1; i < 8; i++ {
		fmt.Printf("%v ", i)
	}
	// for 2
	fmt.Println("\nfor 2")
	j := 1
	for j < 8 {
		fmt.Printf("%v ", j)
		j++
	}
	// for 3
	fmt.Println("\nfor 3")
	k := 1
	for {
		fmt.Printf("%v ", k)
		if k >= 7 {
			break
		}
		k++
	}

	// for 4
	fmt.Println("\nfor 4")
	str := "hello world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("Index %d, Index 对应的字符为 %c \n", i, str[i])
	}

	// for 5
	fmt.Println("for 5")
	for index, value := range str {
		fmt.Printf("Index %d, Index 对应的字符为 %c \n", index, value)
	}

}

func ifDemo() {
	var a int
	fmt.Println("请输入a ：")
	fmt.Scanln(&a)

	if a > 10 { // 可以在 if条件中赋值， 必须加上 {}
		fmt.Println("a 大于10")
		if a < 15 {
			fmt.Println("a 大于 10 小于 15")
		}
	} else if a < 0 {
		fmt.Println("a 小于 0")
	} else {
		fmt.Println("a 小于或等于10,且大于等于0")
	}
}
