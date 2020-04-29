package main
import "fmt"

var num int 
func init() { //初始化函数，默认在main函数前自动被调用
	num = 12
}
func main() {
	fmt.Println(num) // 12
}