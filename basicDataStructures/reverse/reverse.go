package reverse

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func reverseInPlace(s []rune) {
	if len(s) < 2 {
		return
	}

	l := 0
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

	return
}

func reverseWords1(s string) string {
	rune_arr := []rune(s)
	l := 0
	for r := 0; r < len(rune_arr); r++ {
		for r < len(rune_arr) && rune_arr[r] != ' ' {
			r++
		}

		reverseInPlace(rune_arr[l:r])

		for r < len(rune_arr) && rune_arr[r] == ' ' {
			r++
		}
		l = r
	}

	return string(rune_arr)
}

func reverseWords2(s string) string {
	var res strings.Builder
	ss := strings.Split(s, " ")
	for _, w := range ss {
		fmt.Fprintf(&res, "%s ", reverse(w))
	}

	return res.String()[:res.Len()-1]
}
