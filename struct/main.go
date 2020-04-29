package main

import (
	"fmt"
	"strings"
)

type User struct {
	name string
	age  int32
	sex  int32
}

type Person struct {
	firstName string
	lastName  string
}

type Ps Person // 起别名， Ps 代替 Person

func upPerson(person *Person) {
	person.firstName = strings.ToUpper(person.firstName)
	person.lastName = strings.ToUpper(person.lastName)
}

func (person *Person) upPersonPersonFunction() {
	person.firstName = strings.ToUpper(person.firstName)
	person.lastName = strings.ToUpper(person.lastName)
}

func main() {
	func2()
}

func func3() {
	person := Person{"mike", "zhang"}

	person.upPersonPersonFunction()

	fmt.Printf("firstName : %s, lastName: %s", person.firstName, person.lastName)

}

func func2() {
	person := Person{"mike", "zhang"}

	upPerson(&person)

	fmt.Printf("firstName upper is %s, lastName upper is %s\n", person.firstName, person.lastName)

	// person := Ps{"lili", "jane"} 报错 cannot use Ps literal (type Ps) as type Person in assignment
	// 虽然Ps 是Person 的别名，但是Ps 作为独立的一个类型，所以Ps 与 Person 不是同一个类型
	ps := Ps{"lili", "jane"}
	fmt.Printf("firstName upper is %s, lastName upper is %s\n", ps.firstName, ps.lastName)
	// ps.upPersonPersonFunction() 错误的， upPersonPersonFunction是Person 的函数，Person 和 Ps 不是同一个类型

}

func func1() {
	// 测试初始化
	user := &User{}
	userObj := new(User)
	userObj1 := &User{name:"zhangsan", age:12}
	fmt.Println(user) // &{ 0 0}
	fmt.Println(userObj)//&{ 0 0}
	fmt.Println(userObj1)// &{zhangsan 12 0}
	/**
	由上结果所知：
	Go 语言中，未进行显示初始化的变量都会被初始化为该类型的零值，比如 string 零值为空字符串 ， int型为 0
	 */

	var user1 User
	var user2 *User
	user1 = User{"lisi", 25, 2}      //初始化
	user2 = &User{"zhangsan", 24, 1} // 初始化并赋值给结构体指针变量

	fmt.Printf("用户名：%s, 性别： %d，年龄：%d\n", user1.name, user1.sex, user1.age) //用户名：lisi, 性别： 2，年龄：25
	fmt.Printf("用户名：%s, 性别： %d，年龄：%d\n", user2.name, user2.sex, user2.age) //用户名：zhangsan, 性别： 1，年龄：24

	// 测试，“对象”内变量赋值
	user2 = new(User)     // 使用new 给结构体分配内存
	user2.name = "wangwu" // name 赋值
	user2.age = 25
	user2.sex = 2
	fmt.Printf("用户名：%s, 性别： %d，年龄：%d\n", user2.name, user2.sex, user2.age) //用户名：wangwu, 性别： 2，年龄：25
}
