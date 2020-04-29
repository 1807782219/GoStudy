package utils

// 将函数作为参数来调用
func Add(a, b int) int {
	return a + b 
}

func Callback( y int, f func(int, int)int ) int {
	return f(y , 3)
}