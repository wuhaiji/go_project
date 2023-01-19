package test

import "strings"

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	result = make([]string, 0, strings.Count(s, sep)+1)
	for i > -1 {
		term := s[:i]
		if term != "" {
			result = append(result, term)
		}
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Fib2(n int) []int {
	var num1 = 0
	var num2 = 1
	var res = make([]int, 0)
	for i := 0; i <= n; i++ {
		res = append(res, num1)
		res = append(res, num2)
		num1 = num1 + num2
		num2 = num1 + num2
	}
	return res
}
