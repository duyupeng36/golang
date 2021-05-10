package main

import (
	"fmt"
	"net"
	"os"
)

func recvData(conn net.Conn, fileName string) {
	// 按照文件名创建新文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("文件创建出错")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("文件关闭错误")
		}
	}(f)

	// 从 网络中读数据，写入本地文件
	buf := make([]byte, 1024)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("接收文件完成。")
			return
		}
		// 写入本地文件，读多少，写多少。
		_, err = f.Write(buf[:n])
		if err != nil {
			fmt.Println("写入输出错误")
		}
	}
}

func main() {
	// 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("监听套接字创建失败")
		return
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			fmt.Println("监听套接字关闭失败")
			return
		}
	}(listener)

	var conn net.Conn
	conn, err = listener.Accept() // 接收客户端连接
	if err != nil {
		fmt.Println("通信套接字建立失败")
		return
	}
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			fmt.Println("通信套接字关闭失败")
			return
		}
	}(conn)

	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("读取文件名出错")
		return
	}

	fileName := string(buf[:n])

	_, err = conn.Write([]byte("ok"))
	if err != nil {
		fmt.Println("响应数据错误")
		return
	}
	recvData(conn, fileName)
}
