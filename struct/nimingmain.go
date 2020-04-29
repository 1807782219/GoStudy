package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}
func (*innerS) print()  {
	fmt.Println("innerS 内部函数")
}

type outerS struct {
	b      int
	c      float32
	int    //匿名int型字段
	innerS //匿名结构体innerS 型
}
func (out *outerS) print() {
	out.innerS.print()
	// 复写继承的方法，如果没有变化则和innerS 的 print方法一样
}

func main() {
	out := *new(outerS)
	out.b = 12
	out.c = 12.12
	out.int = 233 //对匿名的int型字段赋值
	out.in1 = 32  //对匿名的结构体型字段赋值 = out.innerS.in1 = 32
	out.in2 = 54  //对匿名的结构体型字段赋值  = out.innerS.in2
	out.innerS = innerS{22,33}//对匿名的结构体赋值
	fmt.Println(out)
	fmt.Println(out.in2)
	// 我们通过 out.int 来获取匿名的int型字段的值，所以在一个结构体中对于每一种数据类型只能有一个匿名字段。

	// outerS 结构体初始化
	out1 := outerS{12, 22.12, 34, innerS{34, 35}}
	fmt.Println(out1)

	// 调用“对象”的函数
	out1.print()
	// 调用匿名“对象”的函数
	out1.innerS.print()
}
