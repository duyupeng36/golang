package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func communication(conn net.Conn) {
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			fmt.Printf("服务器与客户端%v连接关闭失败\n", c.RemoteAddr().String())
		}
	}(conn)

	var n int
	var err error
	for {
		// 接收客户端发送过来的数
		var recv []byte
		recv = make([]byte, 1024)
		n, err = conn.Read(recv)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端%v关闭\n", conn.RemoteAddr().String())
				return
			}
			fmt.Printf("读取客户端%v发送的消息错误\n", conn.RemoteAddr().String())
			return
		}

		recvData := string(recv[:n])                                               // 将接收到的数据转为字符串
		fmt.Printf("服务器接收客户端%v接收到的数据为:%v\n", conn.RemoteAddr().String(), recvData) // 输出客户端发送的数据
		recvData = strings.ToUpper(recvData)                                       // 处理接收的数据
		_, err = conn.Write([]byte(recvData))                                      // 将处理好的数据返回给客户端
		if err != nil {
			fmt.Printf("向客户端%v返回数据错误\n", conn.RemoteAddr().String())
			return
		}

		// 如果接收到的数据为quit, 则客户端请求关闭通信
		if recvData == "quit" {
			fmt.Printf("客户端%v请求了连接断开\n", conn.RemoteAddr().String())
			return // 返回
		}

	}
}

func main() {

	// 创建用于监听的socket
	listener, err := net.Listen("tcp", "0.0.0.0:8080") // 创建tcp服务端，监听ip:port 只接收该`ip`发送的连接，服务端运行在8080端口
	// network是小写的tcp或udp
	if err != nil {
		fmt.Printf("net.Listen faild, err=%#v\n", err)
		return
	}
	// 关闭用于监听的socket
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			fmt.Printf("close listener faild, err=%#v\n", err)
		}
	}(listener)

	// 监听客户端链接，并接收链接，生成用于通信的socket
	var conn net.Conn
	for {
		fmt.Printf("服务器等待连接\n")
		conn, err = listener.Accept()
		fmt.Printf("服务器与客户端%v连接成功\n", conn.RemoteAddr().String())
		go communication(conn)
	}
}
