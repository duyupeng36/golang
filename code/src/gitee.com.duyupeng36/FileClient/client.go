package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(conn net.Conn, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("文件关闭失败")
		}
	}(file)

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("文件读取失败")
			}
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("提交数据失败")
			return
		}
	}
}

func main() {
	list := os.Args // 命令行参数
	if len(list) != 2 {
		fmt.Println("没有传递要发送的文件路径，请输入文件路径")
	}
	filePath := list[1] // 获取文件路径

	fileInfo, err := os.Stat(filePath) // 获取文件信息
	if err != nil {
		fmt.Println("文件信获取失败")
		return
	}
	fileName := fileInfo.Name() // 文件名
	fileSize := fileInfo.Size() // 文件大小 字节

	fmt.Printf("%s:%d\n", fileName, fileSize)

	// 发起连接请求
	var conn net.Conn
	conn, err = net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接建立失败")
		return
	}
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			fmt.Println("连接关闭失败")
		}
	}(conn)

	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("文件名发送失败")
		return
	}

	buf := make([]byte, 10)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("读取服务器返回数据失败")
		return
	}
	if "ok" == string(buf[:n]) {
		// 发送文件
		sendFile(conn, filePath)
	}

}
