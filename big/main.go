package main
import (
	"fmt"
	"math"
	"math/big"
)
func main()  {
	im := big.NewInt(math.MaxInt64)
	in := im

	io := big.NewInt(1956)
	ip := big.NewInt(1)

	// Mul 乘法， Add 加法 Div 除法 

	fmt.Println(im)

	ip.Mul(im,in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)

}