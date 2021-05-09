package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080") // 建立tcp连接, ip:port是服务端的地址
	if err != nil {
		fmt.Printf("连接建立错误, err=%#v\n", err)
		return
	}
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			fmt.Printf("连接关闭错误, err=%#v\n", err)
		}
	}(conn)
	reader := bufio.NewReader(os.Stdin) // 创建读取数据对象

	go func(c net.Conn) {
		for {
			var s string
			fmt.Printf("输入数据（quit退出）:")
			s, err = reader.ReadString('\n') // 读取数据
			if err != nil {
				continue // 读取出错，跳过当前循环
			}
			s = strings.TrimSpace(string(s[:])) // 去除\r
			// 发送数据给服务器
			var n int
			n, err = conn.Write([]byte(s)) // 发送数据给服务器
			if err != nil {
				fmt.Printf("数据发送错误, err=%#v\n", err)
				return
			} else {
				fmt.Printf("提交数据的长度:%d\n", n)
			}
			if s == "quit" {
				// 退出客户端
				break
			}
		}
	}(conn)

	for {
		var n int
		// 接收服务端返回的数据
		var recv []byte
		recv = make([]byte, 1024)
		n, err = conn.Read(recv)
		if err != nil {
			fmt.Printf("读取服务端%v返回数据数据失败, err=%#v\n", conn.RemoteAddr().String(), err)
			return
		}
		fmt.Printf("服务端%v返回数据: %#v\n", conn.RemoteAddr().String(), string(recv[:n]))

		// 服务端返回数据为QUIT，则客户端请求了关闭
		if string(recv[:n]) == "QUIT" {
			break
		}

	}
}
