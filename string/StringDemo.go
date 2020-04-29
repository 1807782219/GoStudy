package main

import (
	"fmt"
	"strconv"
)

func main()  {
	
	// stringToBase()
	// zhizhen()
	StringsFunctionDemo()
}

func StringsFunctionDemo() {
	// len() 长度
	str:= "hello"
	fmt.Println(len(str)) // 5
	str = "hello上"
	fmt.Println(len(str)) // 8 Go 采用UTF-8 的字符集，所以一个汉字三个字节
	// 将String 遍历打印 
	for _,value := range str {
		//fmt.Printf("%c", value) // 输出 hello上
		fmt.Printf("%q", value) // 输出 'h''e''l''l''o''上'
	}
	fmt.Println("---------------") 
	// 使用单纯的for循环
	for i:= 0; i<len(str);i++ {
		fmt.Printf("%c", str[i]) //输出 helloä 可以看出 汉字无法输出需要转换
	}

	str2 := []rune(str) // 将无法使用简单for循环遍历不出的汉字，使用rune 转一下，即可正常输出汉字
	for i:= 0; i<len(str2);i++ {
		fmt.Printf("%c", str2[i]) //输出 hello上 
	}

	// String 与 byte 的互转
	byteData := []byte("hello")
	fmt.Printf("\n String 转 byte ： %v\n", byteData)//  String 转 byte ： [104 101 108 108 111]

	// byte 转 String
	str = string([]byte{104,101,108, 108, 111})
	fmt.Printf("\n byte 转 String ： %s\n", str) //byte 转 String ： hello 


}

// 基本数据类型转 String
func baseDataToString() {
	var num1 int = 23
	var num2 float64 = 23.23
	var num3 bool = false
	var mychar byte = 'a'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str 的类型 %T, 和值 %q\n", str, str)
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str 的类型 %T, 和值 %q\n", str, str)
	str = fmt.Sprintf("%t", num3)
	fmt.Printf("str 的类型 %T, 和值 %q\n", str, str)
	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str 的类型 %T, 和值 %q\n", str, str)
	

	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str 的类型 %T, 和值 %q\n", str, str)
	str = strconv.FormatFloat(num2, 'f', 8, 64) //保留8位小数，64表示这个小数是 float64
	fmt.Printf("str 的类型 %T, 和值 %q\n", str, str)
	str =  strconv.FormatBool(num3)
	fmt.Printf("str 的类型 %T, 和值 %q\n", str, str)
	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str 的类型 %T, 和值 %q\n", str, str)
}

// 指针的Demo
func zhizhen() {
	var num int = 10
	var ptr *int = &num

	fmt.Printf("num的地址 %v \n", &num )
	fmt.Printf("num的值 %v \n", num )
	fmt.Printf("指针地址 %v \n", ptr )
	fmt.Printf("指针所指向的值 %v \n", *ptr )
}

// String 数据转基本数据类型
func  stringToBase() {
	num1, num2, num3 := "true", "12","23.23"

	str,err := strconv.ParseBool(num1)
	if err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Println(str)

	str1,err1 := strconv.ParseInt(num2, 10, 64 )
	if err1 != nil {
		fmt.Println(err1)
		return 
	}

	fmt.Println(str1)

	str2,err2 := strconv.ParseFloat(num3, 64)
	if err2 != nil {
		fmt.Println(err2)
		return 
	}

	fmt.Println(str2)


}

