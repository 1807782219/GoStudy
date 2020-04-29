package main
import (
	"fmt"
	"time"
)

func main()  {
	now := time.Now()


	fmt.Println(now) // 2019-11-24 19:31:38.1511152 +0800 CST m=+0.010998301

	fmt.Printf("%d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),now.Day(),now.Hour(), now.Minute(), now.Second()) //2019-11-24 20:35:28
	fmt.Printf(now.Format("2006/01/02 15:04:05\n")) //格式化时间，格式化的形式字符串必须填写这些值的其中一些，形式可以改变 //2019/11/24 20:35:28
	fmt.Printf(now.Format("2006/01/02")) //2019/11/24
}