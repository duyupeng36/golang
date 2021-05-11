package Split

import (
	"strings"
)

func Split(s, sep string) (ret []string) {
	ret = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i >= 0 {
		ret = append(ret, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
