package internal

import (
	"regexp"
	"strings"
)

// Corrects punctuation misplacement
func CorrectPunctuation(text string) string {
	pattern := regexp.MustCompile(`\s+[.!?]{2,3}`)
	text = pattern.ReplaceAllStringFunc(text, func(match string) string {
		i := 0
		for ; i < len(match); i++ {
			if match[i] != ' ' {
				break
			}
		}
		return match[i:]
	})
	pattern = regexp.MustCompile(`\s[.,!?:;]`)
	text = pattern.ReplaceAllStringFunc(text, func(match string) string {
		return string(match[1]) + " "
	})
	pattern = regexp.MustCompile(`'[^']+'`)
	text = pattern.ReplaceAllStringFunc(text, func(match string) string {
		start := 1
		end := len(match) - 2
		for ; match[start] == ' '; start++ {
		}
		for ; match[end] == ' '; end-- {
		}
		return "'" + match[start:end+1] + "'"
	})
	temp := strings.Fields(text)
	var correctedArray []string
	var result string
	for _, s := range temp {
		if s == "" {
			continue
		}
		correctedArray = append(correctedArray, s)
	}
	result = strings.Join(correctedArray, " ")
	return result
}
