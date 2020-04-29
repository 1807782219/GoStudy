package main

import (
	"fmt"
	service "goStudy/service"
)

/*
	多行注释
*/

const PI = 3.14                     // const PI float= 3.14
const STRING string = "stringconst" //const STRING string= "stringconst"
const INTEGER int = 12
const BOOL bool = false
const FU = complex(1, 2) // 1+2i

const beef, two, c = "eat", 2, "veg" //多个
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
const ( //枚举
	Unknown = 0
	Female  = 1
	Male    = 2
)

func dataType() {
	var dateInt int

	dateInt = 3
	println(dateInt + 23)
}

func sayHello() {
	fmt.Println("hello world") // 换行打印
	fmt.Print("aaa\n")         //不换行打印
	fmt.Printf("asdsadas\n")   //不换行打印，格式化打印
	print("aaaaa")             //不换行打印
	println("aaaaa")           //换行打印
}

func main() {
	// sayHello()
	// dataType()
	// sayHello()
	service.SayHello()
	println(Unknown)

}
