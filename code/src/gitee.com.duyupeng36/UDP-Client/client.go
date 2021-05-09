package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 绑定UDP的服务端IP:PORT
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("连接服务器失败")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		var s string
		fmt.Printf("输入要发送的数据(quit退出):")
		s, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("数据读取错误")
			continue
		}

		if strings.TrimSpace(s) == "quit" {
			break
		}

		// 另起一个goroutine用于接收服务端返回的数据
		go func(c net.Conn) {
			var n int
			n, err = c.Write([]byte(s))
			if err != nil {
				fmt.Printf("发送数据失败\n")
				return
			} else {
				fmt.Printf("发送数据的长度为%d\n", n)
			}

			var recv []byte
			recv = make([]byte, 1024)
			n, err = conn.Read(recv)
			if err != nil {
				fmt.Printf("读取服务端数据失败\n")
				return
			}

			fmt.Printf("读取服务端返回的数据为:%v\n", string(recv[:n]))
		}(conn)
	}

}
