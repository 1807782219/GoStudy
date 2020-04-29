package main
import (
    "fmt"
    f "goStudy/utils"
)
// 匿名函数 ，将函数作为返回值
func main() {
    // make an Add2 function, give it a name p2, and call it:
    p2 := Add2()
    fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
    // make a special Adder function, a gets value 3:
    TwoAdder := Adder(2)
    fmt.Printf("The result is: %v\n", TwoAdder(3))
    functionvar := f.FunctionVar // 将外部的匿名函数变量赋值给变量
    functionvar() // 调用外部匿名函数 结果为 外部Utils包下的匿名函数
}

func Add2() func(b int) int {
    return func(b int) int {
        return b + 2
    }
}

func Adder(a int) func(b int) int {
    return func(b int) int {
        return a + b
    }
}

func Test() {
    fmt.Println("test 小写字母开头的函数，在本包可见")
}