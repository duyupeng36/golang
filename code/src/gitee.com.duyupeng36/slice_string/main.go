package main

import "fmt"

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
}
