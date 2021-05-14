package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HttpGet(url string) (html string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("请求发送错误")
		return
	}
	defer func() {
		err := response.Body.Close()
		fmt.Println("关闭主体出错，错误信息为: ", err)
	}()

	result, _ := ioutil.ReadAll(response.Body)
	html += string(result)
	return
}

func main() {
	for i := 0; i < 1; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%8E%8B%E8%80%85%E8%8D%A3%E8%80%80&ie=utf-8&cid=&tab=corearea&pn=" + strconv.Itoa(i*50)
		html := HttpGet(url)
		fmt.Printf("获取到内容|\n%s\n", html)
	}
}
