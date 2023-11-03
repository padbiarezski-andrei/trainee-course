package palindrome

import (
	"regexp"
	"strings"
	"unicode"
)

func Palindrome1(s string) bool {
	//regexp
	var nonAlphanumericRegex = regexp.MustCompile(`[.,\\\/|#!$%\^&\*;:{}=\-_~() ` + "`" + `]+`)
	clearedStr := nonAlphanumericRegex.ReplaceAllString(s, "")

	runeArr := []rune(strings.ToLower(clearedStr))
	if len(runeArr) < 2 {
		return true
	}

	l := 0
	r := len(runeArr) - 1
	for l < r {
		if runeArr[l] != runeArr[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func Palindrome2Helper(r []rune) bool {
	//recursion
	if len(r) <= 1 {
		return true
	}
	if r[0] != r[len(r)-1] {
		return false
	}
	return Palindrome2Helper(r[1 : len(r)-1])
}

func Palindrome2(s string) bool {
	// unicode
	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	return Palindrome2Helper([]rune(clearedStr))
}

func Palindrome3(s string) bool {
	// unicode
	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	runeArr := []rune(clearedStr)
	if len(runeArr) < 2 {
		return true
	}

	l := 0
	r := len(runeArr) - 1
	for l < r {
		if runeArr[l] != runeArr[r] {
			return false
		}
		l++
		r--
	}

	return true
}
