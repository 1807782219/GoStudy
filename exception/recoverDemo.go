package main
import (
	"fmt"
	"errors"
)
// 此例Go 处理异常的方法

func main()  {
	// num := 1

	// defer func ()  {
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println("err = ", err)
	// 	}
	// }()
	// num2 := 0
	// fmt.Println(num / num2)
	// fmt.Println("没有错误继续执行")
	err := writeRight("right")
	if err != nil {
		fmt.Println("出现了错误，错误信息如下：")
		panic(err)
	}
	fmt.Println("没有错误继续执行")
}

func writeRight(str string) error {
	if str == "right" {
		return nil
	}

	return errors.New("输入的不是right")
}