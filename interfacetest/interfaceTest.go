package main

import "fmt"

/**
在Go 语言中，一个类只要实现了接口中的所有函数，则就说这个类实现了接口，没有像Java 必须使用implement 进行强制说明
 */

type IReader interface {
	Read(buf []byte) (n int, err error)
}

type IWriter interface {
	Write(buf []byte) (n int, err error)
}

type ICloser interface {
	Close() error
}

type File struct {
	name string
}

func (file *File) Write(buf []byte) (n int, err error)  {
	fmt.Println("Write")
	return 0, nil
}

func (file *File) Read(buf []byte) (n int, err error)  {
	fmt.Println("Read")
	return 0, nil
}

func (file *File) Close() error  {
	fmt.Println("Close")
	return nil
}
/**
以上代码 File 实现了 IReader、IWriter、ICloser 的接口,Go语言的非嵌入式接口
 */

type A interface {
	A()
	B()
}

type AReal struct {}
func (*AReal) A() {}
func (*AReal) B() {}

type B interface {
	A()
	B()
}
type BReal struct {}
func (*BReal) A() {}
func (*BReal) B() {}

type C interface {
	A()
	B()
	C()
}
type CReal struct {}
func (*CReal) A() {}
func (*CReal) B() {}
func (*CReal) C() {}

func main()  {
	/**
	接口赋值在Go 语言中分为如下两种情况：
	1. 将对象实例赋值给接口，例如fun1() ，要求要被赋值的对象实例实现了接口要求的所有的方法
	2. 将接口A赋值给接口B： 在Go语言中，
		(1)只要两个接口拥有相同的方法列表（次序可忽略），那么他们是等同的，可以相互赋值
		(2)接口赋值并不要求两个接口必须等价，如果接口A 的方法列表是接口B的方法列表的子集，那么接口B的接口对象可以赋值给接口A的接口对象（接口B的方法中有接口A的所有方法）
		例如：fun(2)
	*/
	fun4()
}

func fun1() {
	var file IReader  = new(File)
	file.Read([]byte{1,2,3}) //Read

	var file1 IWriter  = new(File)
	file1.Write([]byte{1,2,3}) //Write

	var file2 ICloser  = new(File)
	file2.Close() //Close
}

/**
接口赋值测试
 */
func fun2() {
	var a A = new(AReal)
	var b B = new(BReal)
	var c C = new(CReal)
	a = b
	b = a
	a = c
	// c = a 编译报错
	b = c
	// c = b 编译报错
}

/**
接口查询测试: 断言。
 */
func fun3() {
	var a A = new(AReal)
	var b B = new(BReal)
	var c C = new(CReal)

	if _, ok := a.(B); ok {  // 检查 a 这个对象实例是否实现了B 的接口
		fmt.Println(ok)
	}else {
		fmt.Println("NO")
	}

	if _, ok := b.(A); ok {  // 检查 b 这个对象实例是否实现了A 的接口
		fmt.Println(ok)
	}else {
		fmt.Println("NO")
	}

	if _, ok := a.(C); ok {  // 检查 a 这个对象实例是否实现了C 的接口
		fmt.Println(ok)
	}else {
		fmt.Println("NO")
	}

	if _, ok := c.(A); ok {  // 检查 c 这个对象实例是否实现了A 的接口
		fmt.Println(ok)
	}else {
		fmt.Println("NO")
	}

	if _, ok := c.(C); ok {  // 检查 c 这个对象实例是否实现了C 的接口
		fmt.Println(ok)
	} else {
		fmt.Println("NO")
	}

/**
以上的结果是：
true
true
NO
true
true

例子：c.(A) 即为一个类型断言，表示对象c，它的类型为CReal，它是否实现了接口A。由于c 对应的CReal类型实现了C接口，C接口定义了方法A、B、C 而接口A定义方法A、B
所以可以类型断言成功，相反如果a.(C) 则断言失败。

*/
}

// 类型查询（断言2）
func fun4()  {
	var a interface{} =  new(CCReal)
	switch a.(type) {
	case A:
		fmt.Println("A")
	case B:
		fmt.Println("B")
	case C:
		fmt.Println("C")
	}
}
// 接口组合： 类似匿名组合，接口也可以进行组合 实现 1 + 1 = 2 的效果

type CC interface {
	A
	C()
	// 这样这个接口就有了 A() , B() C()
}
type CCReal struct {}
func (*CCReal) A() {}
func (*CCReal) B() {}
func (*CCReal) C() {}

//测试接口组合
func fun5() {
	var c C = new(CReal)
	var cc CC  = new(CCReal)
	c = cc
	cc = c
}


