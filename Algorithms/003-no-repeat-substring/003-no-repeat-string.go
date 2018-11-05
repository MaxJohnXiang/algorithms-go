package main

import (
	"fmt"
)

func main() {
	a := "wejklwerjew"
	fmt.Printf("%s\n", lengthOfLongestSubstring3(a))
}

func lengthOfLongestSubstring(s string) int {
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", s[i])
		location[s[i]] = i
	}
	fmt.Println(location)
	return 1
}

func lengthOfLongestSubstring2(s string) int {
	n := len(s)
	i, j := 0, 0
	b := []rune(s)
	m := make(map[rune]int)
	ans := 0
	for i < n && j < n {
		_, ok := m[b[j]]
		if !ok {
			m[b[j]] = 1
			j++
			if j-i > ans {
				ans = j - i
			}
		} else {
			delete(m, b[i])
			i++
		}
	}
	return ans
}

func lengthOfLongestSubstring3(s string) int {
	n := len(s)
	b := []rune(s)
	m := make(map[rune]int)
	res := 0
	for i, j := 0, 0; j < n; j++ {
		_, ok := m[b[j]]
		if ok {
			if m[b[j]] > i {
				i = m[b[j]]
			}
		}
		if j-i+1 > res {
			res = j - i + 1
		}
		m[b[j]] = j + 1
	}
	return res
}
func lengthOfLongestSubstring4(s string) int {
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}
	i, j := 0, 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if location[s[j]] > i {
			i = location[s[j]] + 1
		}
		if j-i+1 > ans {
			ans = j - i + 1
		}
		location[s[j]] = j
	}
}
