package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 使用Get方法获取服务器响应包数据
	//resp, err := http.Get("http://www.baidu.com")
	resp, err := http.Get("https://docs.python.org/zh-cn/3.7/library/http.htmls")

	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()

	// 获取服务器端读到的数据
	fmt.Println("Status = ", resp.Status)         // 状态
	fmt.Println("StatusCode = ", resp.StatusCode) // 状态码
	fmt.Println("Header = ", resp.Header)         // 响应头部
	fmt.Println("Body = ", resp.Body)             // 响应包体

	buf := make([]byte, 4096) // 定义切片缓冲区，存读到的内容
	var result string
	// 获取服务器发送的数据包内容
	for {
		n, err := resp.Body.Read(buf) // 读body中的内容。
		if n == 0 {
			fmt.Println("Body.Read err:", err)
			break
		}
		result += string(buf[:n]) // 累加读到的数据内容
	}
	// 打印从body中读到的所有内容
	fmt.Println("result = ", result)
}
