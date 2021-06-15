package main

import (
	_ "beego/models"
	_ "beego/routers"
	"github.com/astaxie/beego"
)

//ShowNextPage 获取下一页页码
func ShowNextPage(pageIndex int, pageCount int) int {
	if pageIndex == pageCount {
		return pageIndex
	}
	return pageIndex + 1
}

//ShowPrePage 获取上一页页码
func ShowPrePage(pageIndex int) int {
	if pageIndex == 1 {
		return pageIndex
	}
	return pageIndex - 1
}

func main() {
	err := beego.AddFuncMap("next", ShowNextPage)
	if err != nil {
		return
	}
	err = beego.AddFuncMap("pre", ShowPrePage)
	if err != nil {
		return
	}
	beego.Run()
}
