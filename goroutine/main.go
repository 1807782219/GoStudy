package main

import (
	"fmt"
	"sync"
	"runtime"
)

/**
goroutine 协程：
1、Go 语言用于并发的系统资源，特点：与传统的系统级线程和进程相比，协程的最大优势是在于其“轻量级” 可以轻松创建上百万个而不会导致系统资源衰竭而线程和进程通常最多也不过1万个，这也是协程也叫轻量级线程的原因。

2、Go 语言在语言级别上支持轻量级线程，goroutine，Go怨言标准库提供的所有系统调用的操作（包括同步IO操作），都会出让CPU给goroutine，这让协程的切换管理
不依赖系统的线程和进程，也不依赖于CPU的核心数量。

3、在Go 语言中，使用并发时，直接使用go 关键字，该关键字可以让一段代码，在调用的时候会使用一个新的goroutine 进行调用。当被调用的函数返回时，
这个goroutine也就自动结束了
 */
/**
并发通信： 共享数据和消息
1、共享数据是真多个并发单元分别保持对同一个数据的引用，实现对该数据的共享，被共享的数据可能有多种形式，比如内存数据块、磁盘文件、网络数据等，
在实际工程应用中最常用的无疑是内存了，也就是常说的共享内存。
2、消息机制认为每个并发单元是自含的、独立的个体，并且都有自己的变量，但在不同并发单元间这些变量不共享。每个并发单元的输入和输出只有一种，那就是消息。
这有点类似于进程，每个进程不会被其他进程打扰，它只做好自己的工作就像，不同进程间靠消息来通信，他们不会共享内存
Go 语言提供的消息通信机制被称为channel。 “不要通过共享内存来通信，而应该通过通信来共享内存”


 */
func Add(x, y int) {
	fmt.Println(x + y)
}

// 共享内存 测试案例：
var counter int = 0
func Count(lock *sync.Mutex)  {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}
func mainFunc1()  {
	lock := &sync.Mutex{}
	for i:=0;i<=10;i++ {
		go Count(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()

		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}

/**
Go 语言提供的消息通信机制被称为channel.在两个或多个goroutine之间进行传递消息，channel 是进程内的通信方式，所以通过channel传递对象的
过程和调用函数时的参数传递行为比较一致。
声明与定义：与一般的变量声明不同的地方仅仅是在类型之前加了chan 关键字。
声明：var a chan int  或 var m mao[string] chan bool
定义：ch := make(chan int) 声明并定义个一个ch 的channel
在channel 的用法中，最基本就是写入和写出
写入 ch<- value , 向channel 写入数据通常会导致程序阻塞，直到有其他goroutine 从这个channel 中读取数据
写出 value := <-ch， 如果channel 之前没有写入数据，那么从channel 中读取数据也会导致程序阻塞，直到channel中被写入数据为止。
select 用于处理IO问题
语法：
select {
	case ch <- 1:
	...
	case <-ch:
	....
	default:
	...

缓冲机制：用于持续传输大数据场景
在定义的时候指定大小，ch := make(chan int,1024),这样即使没有读取方，写入方也可以一直写到channel中，直到缓冲区被填满

单向channel ：只能用于发送或接收数据。channel本身是同时指出读写的，所谓的channel概念，其实是对channel的一种使用限制
定义：	var ch chan int // 一个正常的channel，支持读写
		var ch1 chan<- int// 只支持写
		var ch2 <-chan int // 只支持读取
初始化：
		ch = make(chan int)
		ch1 = chan<- int(ch)
		ch2 = <-chan int(ch)
即，使用强制类型转换将双通道的变量转为单通道的

关闭channel: close(ch) 判断是否被关闭：
if _,ok := <-ch; !ok {
	fmt.Println("关闭成功")
	if _,ok := <-ch; !ok {
		fmt.Println("关闭成功")
	}
}

出让时间片：runtime.Gosched() 调用该函数后，正在执行的goroutine 主动让出时间片给其他goroutine，让其执行

同步：
	同步锁，Go 语言中的sync 包中提供了两种锁类型，sync.Mutex 和 sync.RWMutex. Mutex 是最简单的一种锁类型，同时也比较暴力，当一个goroutine获得了Mutex后，
	其他 goroutine 就只能乖乖等待，直到该goroutine释放Mutex。RWMutex相对友好写，在读锁占用的情况下，会阻止写，但不阻止读，也就是多个goroutine可以同时获取
	读锁（调用RLock()）；而写锁被占用情况下（调用Lock()），会阻止任何其他goroutine进来，整个锁相当于由该goroutine独占。

	对于这两种锁类型，任何一个Lock() 和 RLock() 均需要保证对应有Unlock() 或RUnlock() 调用与之对应，否则可能导致等待该锁的所有goroutine处于饥饿状态，甚至
	可能导致死锁。
	典型使用模式如下：
	var lock sync.Mutex
	func fun() {
		lock.Lock()
		defer lock.Unlock()

	}

全局唯一性操作：对于全局的角度值需要运行一次的代码，比如全局初始化操作，Go语言提供了一个Once类型来保证全局的唯一性操作

 */
// 通过消息进行控制
func CouterByChannel(ch chan int) {
	ch <- 1
	fmt.Println("CouterByChannel")
}
func mainFunc2() {
	chs := make([]chan int, 10)
	for i:=0; i< 10; i++ {
		chs[i] = make(chan int)
		go CouterByChannel(chs[i])
	}

	for _, ch := range chs{
		<-ch
	}
}

// Once 测试
var a string
var once sync.Once

func setup()  {
	a = "hello world"
}

func doPrint()  {
	once.Do(setup)
	fmt.Println(a)
}

func funmain2()  {
	go doPrint()
	go doPrint()
}

func main() {
	//mainFunc1()
	//mainFunc2()
	//for i:=0; i<10; i++ {
	//	go Add(i, i+1)
	//}
	funmain2()
}
