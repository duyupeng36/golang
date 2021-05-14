package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"os"
	"regexp"
	"strings"
)

var jsonByteSlice = make([]map[string]string, 0, 1000)

// ParasUrl 生成需要访问的url
func ParasUrl(start, end int) (urls []string) {
	var url = "https://www.qiushibaike.com/text/page/%d/"
	urls = make([]string, 0, end-start+1)
	for i := start; i <= end; i++ {
		if i == 1 {
			urls = append(urls, "https://www.qiushibaike.com/text")
		} else {
			urls = append(urls, fmt.Sprintf(url, i))
		}
	}
	return
}

func NewRequest(method, url string, headers, params map[string]string) (request *http.Request, err error) {
	request, err = http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	for k, v := range headers {
		request.Header.Add(k, v)
	}
	urlParams := new(url2.Values)
	for k, v := range params {
		urlParams.Add(k, v)
	}

	request.URL.RawQuery = urlParams.Encode()

	return
}

// InitiateRequest 发起请求获得响应内容
func InitiateRequest(url string, in chan<- string) {
	// 设置请求头
	fmt.Printf("正在爬取url=%s\n", url)
	headers := make(map[string]string, 10)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"
	request, err := NewRequest("GET", url, headers, nil)
	if err != nil {
		return
	}
	var response *http.Response
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	var htmlByte []byte
	htmlByte, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	in <- string(htmlByte)
}

func ParamsHTML(html string, f *os.File) {

	nick := `<h2>(?s:.(?P<nickName>.*?).)</h2>`
	reNick, err := regexp.Compile(nick)
	if err != nil {
		return
	}

	result := reNick.FindAllStringSubmatch(html, -1)
	var nickNameSlice = make([]string, 0, len(result))
	for _, slice := range result {
		nickNameSlice = append(nickNameSlice, slice[1])
	}

	content := `(?s:<div class="content">.<span>.(?s:(.*?)).</span>)`
	reContent, err := regexp.Compile(content)
	if err != nil {
		return
	}
	result = reContent.FindAllStringSubmatch(html, -1)
	var contentSlice = make([]string, 0, len(result))
	for _, v := range result {

		c := strings.TrimSpace(v[1])

		contentSlice = append(contentSlice, c)
	}
	var NickContentMap = make(map[string]string, 100)
	for i := 0; i < len(nickNameSlice); i++ {
		NickContentMap[nickNameSlice[i]] = contentSlice[i]
	}
	jsonByteSlice = append(jsonByteSlice, NickContentMap)
}

func main() {
	ch := make(chan string, 100)
	f, err := os.Create("C:\\Users\\dyp\\Desktop\\golang\\code\\src\\Spider\\FML\\糗事百科段子.json")
	if err != nil {
		return
	}
	var (
		start int
		end   int
	)
	fmt.Print("输入起始页码与结束页码(空格分隔): ")
	fmt.Scanln(&start, &end)
	urls := ParasUrl(start, end)
	for _, v := range urls {
		go InitiateRequest(v, ch)
	}

	num := 0
	for {
		select {
		case html := <-ch:
			{
				num++
				go ParamsHTML(html, f)
				if num == end-start+1 {
					goto A
				}
			}
		}
	}
A:
	b, e := json.Marshal(jsonByteSlice)
	if e != nil {
		fmt.Println("序列化错误")
		return
	}
	_, err = f.Write(b)
	if err != nil {
		fmt.Println("数据写入错误")
		return
	}
	_ = f.Close()
}
