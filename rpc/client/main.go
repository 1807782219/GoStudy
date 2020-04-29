package main

import (
	"net/rpc"
	"log"
	"goStudy/rpc/server"
	"fmt"
)

func main() {
	serverAddress := "127.0.0.1"
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("连接出错-》:", err)
	}

	args := server.Args{10,2}

	// 调用乘法服务
	var reply int
	err = client.Call("Arith.Multipy", args, &reply)
	if err != nil {
		log.Fatal("乘法服务调用出错-》", err)
	}
	fmt.Println(reply)

	// 调用除法余数服务
	var qutoient server.Quotient
	err = client.Call("Arith.Divide", args, &qutoient)
	if err != nil {
		log.Fatal("除法服务调用出错-》", err)
	}
	fmt.Println(qutoient.Quo)
	fmt.Println(qutoient.Rem)
}
