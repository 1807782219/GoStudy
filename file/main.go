package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"io/ioutil"
)

func main()  {
	testReadAllDataFromFile()
}

/**
适用场景： 要读取的文件比较大
 */
func testReadBigDataFromFile()  {
	file, err := os.Open("F:/goworkspace/src/goStudy/file/1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 读取
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			//已经到达文件的末尾
			break
		}
		fmt.Println(str)
	}
}

/**
应用场景： 读取数据量小的文件
 */
func testReadAllDataFromFile() {
	filePath := "F:/goworkspace/src/goStudy/file/1.txt"
	arr, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(arr))
}