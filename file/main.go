package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"io/ioutil"
	"log"
)

func main()  {
	fromFilePath := "F:/goworkspace/src/goStudy/file/2.txt"
	toFilePath := "F:/goworkspace/src/goStudy/file/3.txt"
	_,  err := testCopyFile(fromFilePath, toFilePath)
	if err != nil {
		fmt.Println("拷贝出错》", err)
	} else {
		fmt.Println("拷贝成功")
	}
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

/**
写文件：
# 文档： https://studygolang.com/static/pkgdoc/pkg/os.htm#OpenFile
func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
os.OpenFile 是一个更一般行的文件打开函数： name 文件名，flag 打开的方式 ， perm 权限控制（仅适用于linux、unix、MACOS 等，不适合Windows）
flag:
const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
perm : 例如 0666等
 */
// 向一个不存在的文件中，写入五行 hello world
func testWriteToNoExistFile() {
	filePath := "F:/goworkspace/src/goStudy/file/2.txt"
	file, err := os.OpenFile(filePath,  os.O_CREATE | os.O_WRONLY , 0666)
	if err != nil {
		log.Fatal("打开文件出错》", err)
	}

	// 关闭文件
	defer file.Close()

	writer := bufio.NewWriter(file)
	str := "hello world !\n"
	for i:=0; i<5; i++  {
		writer.WriteString(str)
	}
	//因为writer 是带缓存的，因此在使用WriteString 函数时，其实内存先写入缓存中，所以如果需要将数据保存到磁盘需要调用 Flush() 进行刷新，
	// 否则文件中不会有数据
	writer.Flush()

}

// 向一个已经存在的文件中，替换文件中内容
func testWriteToExistFile() {
	filePath := "F:/goworkspace/src/goStudy/file/2.txt"
	file, err := os.OpenFile(filePath,  os.O_WRONLY | os.O_TRUNC , 0666)
	if err != nil {
		log.Fatal("打开文件出错》", err)
	}

	// 关闭文件
	defer file.Close()

	writer := bufio.NewWriter(file)
	str := "你好，世界\n"
	for i:=0; i<5; i++  {
		writer.WriteString(str)
	}
	//因为writer 是带缓存的，因此在使用WriteString 函数时，其实内存先写入缓存中，所以如果需要将数据保存到磁盘需要调用 Flush() 进行刷新，
	// 否则文件中不会有数据
	writer.Flush()

}

// 向一个已经存在的文件中，替换文件中内容
func testAppendWriteToExistFile() {
	filePath := "F:/goworkspace/src/goStudy/file/2.txt"
	file, err := os.OpenFile(filePath,  os.O_WRONLY | os.O_APPEND , 0666)
	if err != nil {
		log.Fatal("打开文件出错》", err)
	}

	// 关闭文件
	defer file.Close()

	writer := bufio.NewWriter(file)
	str := "hello world\n"
	for i:=0; i<5; i++  {
		writer.WriteString(str)
	}
	//因为writer 是带缓存的，因此在使用WriteString 函数时，其实内存先写入缓存中，所以如果需要将数据保存到磁盘需要调用 Flush() 进行刷新，
	// 否则文件中不会有数据
	writer.Flush()

}

// 先读取文件内容，后追加写入数据
func testReadFileThenAppendWriteToExistFile() {
	filePath := "F:/goworkspace/src/goStudy/file/2.txt"
	file, err := os.OpenFile(filePath,  os.O_RDWR | os.O_APPEND , 0666)
	if err != nil {
		log.Fatal("打开文件出错》", err)
	}

	// 关闭文件
	defer file.Close()

	//读取数据
	reader := bufio.NewReader(file)
	for {
		str, er := reader.ReadString('\n')
		if er == io.EOF {
			fmt.Println("文件已到末尾 》", er)
			break
		}
		fmt.Println(str)

	}

	// 开始写入数据
	writer := bufio.NewWriter(file)
	str := "再次追加\n"
	for i:=0; i<5; i++  {
		writer.WriteString(str)
	}
	//因为writer 是带缓存的，因此在使用WriteString 函数时，其实内存先写入缓存中，所以如果需要将数据保存到磁盘需要调用 Flush() 进行刷新，
	// 否则文件中不会有数据
	writer.Flush()

}

/**
文件拷贝 :
可以使用到 io 包下的Copy函数
func Copy(dst Writer, src Reader) (written int64, err error)
  */
func testCopyFile(fromFile string, toFile string) (written int64, err error) {
	fileFrom,err := os.Open(fromFile)
	if err != nil {
		fmt.Println("文件打开失败",err)
		return 0, err
	}
	defer fileFrom.Close()
	reader := bufio.NewReader(fileFrom)

	fileTo,err := os.OpenFile(toFile, os.O_CREATE | os.O_WRONLY , 0666)
	if err != nil {
		fmt.Println("文件打开失败",err)
		return 0, err
	}

	defer fileTo.Close()

	writer := bufio.NewWriter(fileTo)
	defer writer.Flush() // 这里必须要有，因为博主在第一次执行的时候发现文件有了，但是文件中的内容没有，虽然输出拷贝成功。 查看Copy函数的源码发现没有使用Flush() 函数。
	//所以我们在拷贝完成后需要自己手动刷新一下。
	return io.Copy(writer, reader)
}


















