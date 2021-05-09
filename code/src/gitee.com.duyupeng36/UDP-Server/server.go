package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var conn *net.UDPConn
	// 创建UDP地址
	serverAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8080")
	if err != nil {
		fmt.Printf("地址创建失败\n")
		return
	}
	// 创建用于通信UDP socket
	conn, err = net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Printf("UDP socket创建失败\n")
		return
	}
	defer func(c *net.UDPConn) {
		err := c.Close()
		if err != nil {
			fmt.Printf("连接关闭失败\n")
		}
	}(conn)

	// 循环的一处理客户端的
	for {
		var recv []byte
		recv = make([]byte, 1024)
		var n int
		var clientAddr *net.UDPAddr
		fmt.Printf("等待读取客户端发送来的数据\n")
		n, clientAddr, err = conn.ReadFromUDP(recv) // 阻塞的
		if err != nil {
			fmt.Printf("数据读取失败\n")
		} else {
			fmt.Printf("读取数据的长度: %d\n", n)
		}
		fmt.Printf("客户端%v发送来的数据为: %v\n", clientAddr, string(recv[:n]))

		// 另起一个goroutine用于返回数据给客户端
		go func(c *net.UDPConn) {
			now := time.Now().String()

			_, err = c.WriteToUDP([]byte(string(recv[:n])+now), clientAddr)
			if err != nil {
				fmt.Printf("WriteToUDP err=%v\n", err)
			}
		}(conn)
	}
}
