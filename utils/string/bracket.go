package utils

import (
	"regexp"
	"strings"
)

func FindFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
	return ""
}

//FindFirstStringInBracketV2 find and get string in bracket
func FindFirstStringInBracketV2(str string) string {
	//validate double bracket
	str = regexp.MustCompile(`(?m)([(])+`).ReplaceAllString(str, "(")
	str = regexp.MustCompile(`(?m)([)])+`).ReplaceAllString(str, ")")

	//first index bracket
	first := strings.IndexByte(str, '(')
	if first < 0 {
		return ""
	}

	//last index bracket
	last := strings.IndexByte(str, ')')
	if last < 0 {
		return ""
	}

	//get string first bracket
	return str[first+1 : last]
}
