package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"os"
	"regexp"
	"strconv"
)

type move struct {
	Name  string
	Score float64
	Num   int
}

var moveSlice = make([]move, 0, 250)
var err error // 错误
func ParasHTML(html string) {

	var result *regexp.Regexp
	result, err = regexp.Compile(`(?s:<img width="100" alt="(?P<moveName>.*?)" src=".*?" class="">.*?<span class="rating_num" property="v:average">(?P<score>.*?)</span>.*?<span>(?P<num>\d+)人评价</span>)`)
	r := result.FindAllStringSubmatch(html, -1)

	for _, one := range r {
		var score float64
		var num int
		name := one[1]
		score, err = strconv.ParseFloat(one[2], 64)
		num, err = strconv.Atoi(one[3])

		moveSlice = append(moveSlice, move{
			Name:  name,
			Score: score,
			Num:   num,
		})
	}
}

func main() {
	var request *http.Request            // 请求对象
	var respByte []byte                  // 读取出的数据
	var response *http.Response          // 响应对象
	var respByteChan = make(chan string) // 数据通道
	var quit = make(chan bool)           // 用于退出

	headers := make(map[string]string, 10)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"

	url := "https://movie.douban.com/top250"

	fmt.Print("输入需要爬取的页数起止页数(空格分隔): ")
	var start int
	var end int

	_, err = fmt.Scanln(&start, &end)
	if err != nil {
		fmt.Printf("读取输入失败，失败原因: %v\n", err)
		return
	}
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("请求对象创建失败，失败原因: %v\n", err)
		return
	}
	// 添加请求头
	for k, v := range headers {
		request.Header.Add(k, v)
	}

	// 创建一个goroutine用于提起数据
	go func() {
		for {
			select {
			case html := <-respByteChan:
				ParasHTML(html)
			case <-quit:
			}
		}

	}()

	for i := start; i <= end; i++ {
		// 创建请求url参数
		params := make(url2.Values, 2)
		params.Add("start", strconv.Itoa((i-1)*25))
		params.Add("filter", "")
		request.URL.RawQuery = params.Encode() // 添加url参数

		response, err = http.DefaultClient.Do(request)
		if err != nil {
			fmt.Printf("请求失败，失败原因: %v\n", err)
			return
		}

		respByte, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("读取数据错误，错误原因: %v\n", err)
			return
		}
		respByteChan <- string(respByte)
	}
	quit <- true

	fmt.Println(moveSlice)
	var movesByte []byte
	movesByte, err = json.Marshal(moveSlice)
	if err != nil {
		fmt.Printf("序列化错误，错误原因:%v\n", err)
	}
	var f *os.File
	f, err = os.Create("C:\\Users\\dyp\\Desktop\\golang\\code\\src\\Spider\\DoubanMovie\\DoubanMove.json")
	if err != nil {
		fmt.Printf("创建文件错误，错误原因:%v\n", err)
	}
	_, err = f.Write(movesByte)
	if err != nil {
		fmt.Printf("写入文件错误，错误原因:%v\n", err)
	}
}
