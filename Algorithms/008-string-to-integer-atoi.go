package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	x := stringToIntegerAtoi("-433x1")
	fmt.Printf("convert result is %s\n", x)
}

func stringToIntegerAtoi(s string) int {
	//首先去掉空格
	s = strings.TrimSpace(s)
	//判断正负
	sign, x := getSign(s)
	x = trim(x)
	k := convert(sign, x)
	return sign * k
}

func trim(s string) string {
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return s[:i]
		}
	}
	return s
}

func getSign(s string) (int, string) {
	sign := 1
	switch s[0] {
	case '-':
		s = s[1:]
		sign = -1
	case '+':
		s = s[1:]
	default:
	}
	return sign, s
}

func convert(sign int, s string) int {
	base := 1
	res := 0
	isOver := false
	for i := len(s) - 1; i >= 0; i-- {
		res, isOver = isOverFlow(res + ((int(s[i]) - 48) * base))
		fmt.Printf("res is %d\n", res)
		if isOver {
			return res
		}
		base = base * 10
	}
	return res
}

func isOverFlow(i int) (int, bool) {
	switch {
	case i > math.MaxInt32:
		return math.MaxInt32, true
	case i < math.MinInt32:
		return math.MinInt32, true
	default:
		return i, false
	}
}
