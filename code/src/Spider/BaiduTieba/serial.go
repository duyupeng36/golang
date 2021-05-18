package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"os"
	"strconv"
)

var url = "https://tieba.baidu.com/f"

var ch = make(chan map[int][]byte, 100)

func initiateRequest(req *http.Request, i int) {
	fmt.Printf("正在请求第%d页\n", i)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("请求发送失败\n")
		return
	}
	ret, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("数据读取失败\n")
	}
	var m = map[int][]byte{i: ret}
	ch <- m
}

func main() {
	fmt.Print("输入需要抓取的页数:")
	pageNum := new(int)
	fmt.Scan(pageNum)
	for i := 0; i < *pageNum; i++ {
		req, _ := http.NewRequest(http.MethodGet, url, nil) // 创建请求对象
		// 添加url参数
		params := make(url2.Values)
		params.Add("kw", "王者荣耀")
		params.Add("ie", "utf-8")
		params.Add("cid", "")
		params.Add("tab", "corearea")
		params.Add("pn", strconv.Itoa(i*50))
		req.URL.RawQuery = params.Encode()

		// 添加请求头
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
		go initiateRequest(req, i)
	}

	num := 0
	for {
		select {
		case pageMap := <-ch:
			num++

			for k, v := range pageMap {
				f, _ := os.Create(fmt.Sprintf("第 %d 页.html", k))
				writer := bufio.NewWriter(f)
				_, err := writer.Write(v)
				if err != nil {
					fmt.Printf("写入第 %d 页数据出错, 错误原因: %v\n", k, err)
				}
			}
			fmt.Printf("第%d页数保存完成\n", num)
			if num == *pageNum {
				fmt.Println("所有数据下载完成")
				return
			}
		}
	}
}