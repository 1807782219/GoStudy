package main

import (
	"goStudy/rpc/server"
	"net/rpc"
	"net"
	"log"
	"net/http"
	"fmt"
)

/**
RPC(远程过程调用) 是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络细节的应用程序通信协议。RPC 协议构建与TCP或UDP，或者是HTTP 之上，
允许开发者直接调用另一台计算机上的程序，而开发者无需额外的为这个调用过程编写网络通信相关代码，使得开发包括网络分布式程序在内的应用程序更加容易。
net/rpc 包允许RPC 客户端程序通过网络或其他IO链接链接调用一个远程对象的公开方法（必须是大写字母开头，可外部调用），在RPC服务器，可将一个对象注册为可访问的服务，
之后该对象的公开方法就能够以远程的方式提供访问。一个RPC 服务端，可以注册多个不同类型的对象，但不允许注册同一类型的多个对象
服务端，一个对象只有满足如下条件条件的方法，才能被RPC服务器设置为可供远程访问：
1、 必须是在对象外部可公开调用的方法（首字母大写）
2、 必须有两个参数，且参数的类型都必须是包外部可以访问的类型或者Go内建支持的类型
3、 第二个参数必须是一个指针
4、 方法必须返回一个error类型的值
如： func (t *T) MethodName(argType T1, replyType *T2) error ；
1、 类型T、T1 和 T2 默认会使用 Go 内置的 encoding/gob 包进行编码解码
2、 方法的第一个参数，表示由RPC客户端传入的参数，第二个参数表示要返回给RPC客户端的接口，该方法最后返回一个error 类型的值。
RPC 服务端可以通过调用rpc.ServerConn() 处理单个连接请求，多数情况下，通过TCP或是HTTP在某个网络地址上进行监听来创建该服务器是个不错的选择。
在RPC 客户端，GO的net/rpc 包提供了遍历的rpc.Dial() 和 rpc.DialHTTP() 方法来与指定的RPC 服务端建立连接。在建立连接之后，GO net/rpc 包允许我们
使用同步或者异步的方式接收RPC 服务器的处理结果。调用RPC客户端的处理结果之后，才开始继续执行后面的程序，当调用RPC客户端的Go方法时，则可以进行异步处理，
RPC客户端程序无需等待服务端的结果即可执行后面的程序，而当接收到RPC服务端的处理结果时，在对其进行相应的处理。无论是调用RPC客户端的Call() 或者是 Go() 方法。都必须
指定要调用的服务及其方法名称以及一个客户端传入参数的引用，还有一个用户接收处理结果参数的指针。



 */
func main() {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l,err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal("listen error ....", err)
	}
	defer l.Close()
	fmt.Println("kais1")
	e := http.Serve(l, nil)
	if e != nil {
		log.Fatal("Serve error ....", err)
	}
}
