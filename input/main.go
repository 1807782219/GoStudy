package main

import "fmt"
func main () {
	var name string 
	var age int 
	var sal float32
	var isPass bool

	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过：")
	fmt.Scanln(&isPass)

	fmt.Printf("%v , %v, %v , %v \n", name, age, sal, isPass)


	fmt.Println("==========第二种 =================0")
	fmt.Scanf("%s %d %f %t", &name , &age ,&sal, &isPass)
	fmt.Printf("%v , %v, %v , %v ", name, age, sal, isPass)
	
}