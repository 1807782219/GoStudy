package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func (tI *TwoInts) Add() int {
	return tI.a + tI.b
}
func (tI *TwoInts) AddPlus(params ...int) int {

	sum := tI.a + tI.b

	for _, v := range params {
		sum = sum + v
	}

	return sum
}
func (tI *TwoInts) fun2() {
	tI.a = 2 * tI.a
}
func main() {
	tI := TwoInts{12, 12}
	sum := tI.Add()
	fmt.Println(sum)

	sum = tI.AddPlus(1, 2, 3, 4) //34
	fmt.Println(sum)

	tI.fun2()
	fmt.Println(tI.a)
}
