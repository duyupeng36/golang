package main

import (
	"fmt"
)

func BruteForce(dest, patten string) int {
	d := []rune(dest)
	p := []rune(patten)
	var i, j int

	// 两个串都没完成匹配继续循环
	for i < len(d) && j < len(p) {
		if d[i] == p[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j >= len(p) {
		return i - len(p)
	} else {
		return -1
	}
}

func GetNext(patten string) (next []int) {
	p := []rune(patten)
	next = make([]int, len(p))
	j := 0
	k := -1
	next[0] = -1

	for j < len(p)-1 {
		if k == -1 || p[j] == p[k] {
			j++
			k++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return
}

func GetNextval(patten string) (next []int) {
	p := []rune(patten)
	next = make([]int, len(p))
	j := 0
	k := -1
	next[0] = -1
	for j < len(p)-1 {
		if k == -1 || p[j] == p[k] {
			j++
			k++
			if p[j] != p[k] {
				next[j] = k
			} else {
				next[j] = next[k]
			}
		} else {
			k = next[k]
		}
	}
	return
}

func KMP(dest, patten string) int {
	next := GetNextval(patten)
	i := 0
	j := 0
	d := []rune(dest)
	p := []rune(patten)
	fmt.Println(next)
	for i < len(d) && j < len(p) {
		if j == -1 || d[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j >= len(p) {
		return i - len(p)
	} else {
		return -1
	}
}
func main() {

	fmt.Println(KMP("aaaabbaa", "aabba"))

}
