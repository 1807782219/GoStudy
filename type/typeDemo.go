package main
import (
	"fmt"
	"unsafe"
)
type TZ int
func main() {
	var a, b TZ = 2,3
	c := a +b 
	fmt.Println( c )
	// unsafe.Sizeof(c) 是unsafe 包中一个函数，Sizeof 是返回一个变量的字节数
	fmt.Printf("c 的类型 %T, 字节数是 %d \n", c, unsafe.Sizeof(c))

	floatnum1 := 3.14159
	floatnum2 := .312131
	floatnum3 := 3.1419e2

	fmt.Printf("floatnum1: %v  , floatnum2: %v , floatnum3: %v \n" , floatnum1,floatnum2,floatnum3)

	char3 := 1
	fmt.Println(char3) // 码值
	fmt.Printf("char3 的字符形态 %c \n", char3)
	var char byte = '1'
	fmt.Println(char) // 码值
	fmt.Printf("char 的字符形态 %c \n", char)
	var char2 int = '我'
	fmt.Println(char2) // 码值
	fmt.Printf("char2 的字符形态 %c \n", char2)

	str := `
		package main
		func main() {
			fmt.Println("hello world!")
		}
	`

	fmt.Println(str)
	dataZhuan() 

}


func dataZhuan() {
	var num int32 = 256
	var num1 float32 = float32(num)
	var num2 int8 = int8(num)
	fmt.Println(num1)
	fmt.Println(num2)
}




