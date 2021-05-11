package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"
	"time"
)

var err error

// Client 用于描述客户端的结构体
type Client struct {
	C    chan string // 通道用于给客户端发送消息
	Name string      // 客户端的名称
	Addr string      // 客户端的地址
}

// OnLineMap 保存在线用户
var OnLineMap map[string]*Client = make(map[string]*Client, 10000)

// Message 保存用户发送来的消息
var Message chan string = make(chan string, 10000)

// NewClient 创建客户端对象
func NewClient(name string, addr string) *Client {
	return &Client{
		C:    make(chan string),
		Name: name,
		Addr: addr,
	}
}

// Manager 负责在线用户遍历，用户消息广播发送。需要与HandleConnect协程及用户子协程协作完成。
// 管理OnLineMap和Message通道
func Manager() {
	// 循环读取Message通道中的消息
	for {
		msg := <-Message
		for _, client := range OnLineMap {
			client.C <- msg
		}
	}
}

func writeMsgToClient(conn net.Conn, client *Client) {
	// 读取用户通道中的消息，如果有消息向当前用户写入消息
	for msg := range client.C {
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("向客户端%v发送消息失败, 错误信息为: %v\n", client.Addr, err)
		}
	}

}

// MakeMassage 创建消息
func MakeMassage(client *Client, message string) (msg string) {
	msg = fmt.Sprintf("[%s:%s]: %s\n", client.Addr, client.Name, message)
	return
}

// HandleConnect 处理客户端连接
func HandleConnect(conn net.Conn) {
	defer func(c net.Conn) {
		err = c.Close()
		if err != nil {
			fmt.Println("关闭通信套接字错误，原因为: ", err)
		}
	}(conn)

	// 创建登录客户端对象
	name := conn.RemoteAddr().String()
	addr := conn.RemoteAddr().String()
	client := NewClient(name, addr)
	// 将客户端对象保存在OnLineMap中
	OnLineMap[addr] = client

	// 创建一个用于发送消息的goroutine
	go writeMsgToClient(conn, client) // 本函数结束后，该goroutine不会被结束
	// 向全局通道中写入客户端上线信息
	Message <- MakeMassage(client, "login")

	isQuit := make(chan bool) // 用户是否退出的通道

	isActivate := make(chan bool) // 用户是否活跃的通道
	// 创建一个用于接收客户端输入的goroutine
	go func() {

		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Printf("检查到客户端[%s]退出\n", client.Name)
				isQuit <- true
				return
			}
			if err != nil {
				fmt.Println("读取信息错误")
				return
			}

			msg := string(buf[:n])
			msg = strings.TrimSpace(msg)
			switch {
			// 查询在线用户能
			case strings.HasPrefix(msg, "who"):
				_, err = conn.Write([]byte("On Line User List:\n"))
				for _, value := range OnLineMap {
					info := fmt.Sprintf("addr=%s : username=%s\n", value.Addr, value.Name)
					_, err = conn.Write([]byte(info))
					if err != nil {
						fmt.Println("消息通知失败, 失败原因: ", err)
					}
				}
			// 修改用户名
			case strings.HasPrefix(msg, "rename"):

				name := strings.Split(msg, "|")[1]
				client.Name = name
				_, err = conn.Write([]byte("用户名修改成功\n"))
				if err != nil {
					fmt.Println("消息通知失败, 失败原因: ", err)
				}
			// 退出聊天室
			case strings.HasPrefix(msg, "exit"):
				fmt.Printf("检查到客户端[%s]退出\n", client.Name)
				isQuit <- true
				return
			default:

				Message <- MakeMassage(client, msg)
			}

			isActivate <- true // 执行完成整个for都活跃
		}
	}()

	for {
		// 监听isQuite的数据流动
		select {

		// 用户退出
		case <-isQuit:
			close(client.C)                // 关闭客户端通信的通道，结束writeMsgToClient goroutine
			delete(OnLineMap, client.Addr) // 删除用户
			Message <- MakeMassage(client, "已退出")
			runtime.Goexit() // 关闭与当前客户端相关的goroutine
		// 超时强制退出
		case <-time.After(time.Minute):
			delete(OnLineMap, client.Addr) // 删除用户
			Message <- MakeMassage(client, "已被强制退出")
			runtime.Goexit() // 关闭与当前客户端相关的goroutine
		case <-isActivate:
			fmt.Println("用户活越，计时重置")
		}
	}

}

// 负责监听、接收用户（客户端）连接请求，建立通信关系。同时启动相应的协程处理任务。
func main() {
	var listener net.Listener
	listener, err = net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("监听套接字创建失败，失败原因为: ", err)
		return
	}
	defer func(l net.Listener) {
		err = l.Close()
		if err != nil {
			fmt.Println("")
		}
	}(listener)
	// 创建用于管理发送消息的goroutine
	go Manager()
	for {
		fmt.Println("等待客户端连接")
		var conn net.Conn
		conn, err = listener.Accept()
		if err != nil {
			fmt.Println("创建通信套接字失败，失败原因为: ", err)
		}
		go HandleConnect(conn)
	}
}
