package palindrome

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Palindrome1(s string) bool {
	//regexp
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	nonAlphanumericRegex.ReplaceAllString(s, "")
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

func Palindrome2(s string) bool {
	panic("not implemented")
	//need to rewrite
	len := utf8.RuneCountInString(s)
	for len > 1 {
		l, sizeL := utf8.DecodeRuneInString(s)
		r, sizeR := utf8.DecodeLastRuneInString(s)
		if l != r {
			return false
		}
		len -= sizeL + sizeR
	}

	return true
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
