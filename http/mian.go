package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
)

/**
Http 编程：Go语言标准库内建提供了net/http包，涵盖了HTTP客户端和服务端的具体实现。
Http 客户端：Go 内置的net/http 包提供了最简洁的http 客户端实现，我们无需借助第三方网络通信库（例如libcurl）就可以直接使用http用得最多的
			GET和POST 方式请求数据。
func (c *Client) Get(url string) (r *Response, err error)
请求一个资源，值需要调用http.Get() 方法即可，
func (c *Client) Post(url string, bodyType string, body io.Reader) (r *Response, err error)
func (c *Client) PostForm(url string, data url.Values) (r *Response, err error)
func (c *Client) Head(url string) (r *Response, err error)
func (c *Client) Do(req *Request) (resp *Response, err error)
 */

func testGet()  {
	res,err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("有错误")
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
	//fmt.Printf("%s", os.Stdout)
}
func main() {
	testGet()
}
