package main

import "fmt"
type IFly interface {
	Fly()
}
type Bird struct {

}

const (
	c0 = iota//0
	c1 = iota//1
	c2 //2
	c3 = iota//3
)
func (b *Bird) Fly() {
	fmt.Print("鸟飞行")
}

func main() {
	var fly IFly = new(Bird)
	fly.Fly()

	a := 11.22
	b := 12
	c := "12345aasdf"

	fmt.Println("test, ", a)

	fmt.Printf("%f, %d, %s\n", a,b,c)
	fmt.Printf("%v, %v, %v\n", a,b,c)

	num1, num2 := reverse(1,2)
	fmt.Printf("%v, %v \n", num1,num2)

	for _, char := range c {
		fmt.Printf("%c" , char)
	}
	fmt.Println("\n")
	var array []int
	array = []int{1,2,3}
	for _, arr := range array {
		fmt.Println(arr)
	}

}

func reverse(num1, num2 int64) (int64, int64) {
	if num1 < num2 {
		num1, num2 = num2, num1 // == t = i; i = j; j = t;
	}

	return num1, num2
}