package main

import (
	"fmt"
	"strings"
	"unicode"
)

func noSame(data []string) []string {
	out := data[:1]
A:
	for _, word := range data {
		for _, w := range out {
			if w == word {
				continue A
			}
		}
		out = append(out, word)
	}
	return out
}

func remove(data []int, index int) []int {
	data = append(data[:index], data[index+1:]...)
	return data
}

func main() {
	data := []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	afterData := noSame(data)
	fmt.Println(afterData)

	a := []int{5, 6, 7, 8, 9}
	a = remove(a, 2)

	fmt.Println(a)

	s1 := "hello沙河"
	// 1. 依次获取每个字符
	// 2. 判断是否为汉字
	// 3. 统计汉字出现的次数
	count := 0

	for _, v := range s1 {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println(count)

	s2 := "how do you do ?"
	m1 := make(map[string]int, 10)
	for _, word := range strings.Fields(s2) {
		_, exist := m1[word]
		if exist {
			m1[word] += 1
		} else {
			m1[word] = 1
		}
	}
	fmt.Println(m1)

	s3 := "上海自来水来自海上"
	r := []rune(s3)
	var i = 0
	var j = len(r) - 1
	var isOk = false
	for ; i != j; i++ {
		if r[i] == r[j] {
			isOk = true
		} else {
			isOk = false
			break
		}
		j--
	}

	if isOk {
		fmt.Println("是回文")
	} else {
		fmt.Println("不是回文")
	}
}
