package main

import (
	"fmt"
	"os"
)

func main() {

	// 创建文件，底层调用os.OpenFile("a.txt",  O_RDWR|O_CREATE|O_TRUNC, 0666) 任何都可以操作该文件
	//fp, err := os.Create("a.txt")
	//fp, err := os.OpenFile("a.txt", os.O_RDWR, 0666)
	//if err != nil {
	//	fmt.Println("文件创建失败")
	//	return
	//}
	//// 延时操作，用于关闭文件
	//defer func(fp *os.File) {
	//	err := fp.Close()
	//	if err != nil {
	//		fmt.Println("文件关闭失败")
	//	} else {
	//		fmt.Println("文件关闭")
	//	}
	//}(fp)
	//
	//fmt.Println("文件打开成功")

	//fp.WriteString("hello world")  // 写入字符串 \r\n Windows中的换行符
	//fp.WriteString("你好")
	//fp.Write([]byte{'a', 'b', 'c', 'd', 'e','f'})  // 写入字符切片
	////fp.WriteAt([]byte{'a', 'b', 'c', 'd', 'e'}, 0)  // 指定位置写入文件
	//fp.Seek(-6, io.SeekEnd)  // 移动光标
	//fp.WriteString("移动")  // 在光标后插入数据

	//buf := make([]byte, 100)
	//_, err = fp.Read(buf)
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(string(buf))
	////fmt.Printf("%s\n", buf)

	// 创建缓冲区
	//r := bufio.NewReader(fp)
	//// 循环读取每一行
	//for {
	//	buf, err := r.ReadBytes('\n')
	//	if err != nil {
	//		if err == io.EOF {
	//			fmt.Println("文件读取完成")
	//			break
	//		}
	//
	//		fmt.Println("文件读取错误")
	//	}
	//	fmt.Println(string(buf))
	//}

	//buf := make([]byte, 30)
	//for {
	//	n, err := fp.Read(buf)
	//
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(buf[:n]))
	//}

	list := os.Args // 字符串切片，保存命令行传递的参数
	//fmt.Println(list)
	if len(list) < 2 {
		fmt.Printf("启动格式错误，必须使用: file.exe 参数1 参数2\n")
		return
	}

	name := list[1]

	fileInfo, err := os.Stat(name)
	if err != nil {
		fmt.Printf("获取文件信息错误\n")
	}
	fmt.Println(fileInfo)
}
