package internal

import (
	"strconv"
	"strings"
)

func EditKeywords(input string) string {
	text := strings.Fields(input)     // Разбиваем входную строку на слова и заносим в массив
	temp := make([]string, len(text)) //
	var correctedArray []string
	var result string
	for i := 0; i < len(text); i++ {
		s := text[i]
		switch s {
		case "a", "an", "A", "An":
			correctArticle := articleChecker(text[i-1], text[i+1])
			temp[i] = correctArticle
		case "(cap)":
			temp[i-1] = strings.Title(temp[i-1])
		case "(cap),":
			temp[i-1] = strings.Title(temp[i-1]) + ","
		case "(up)":
			temp[i-1] = strings.ToUpper(temp[i-1])
		case "(up),":
			temp[i-1] = strings.ToUpper(temp[i-1]) + ","
		case "(low)":
			temp[i-1] = strings.ToLower(temp[i-1])
		case "(low),":
			temp[i-1] = strings.ToLower(temp[i-1]) + ","
		case "(cap,":
			count, comma, err := getNrepeats(text[i+1])
			errCheck(err)
			for j := 1; j <= count && i-j >= 0; j++ {
				if temp[i-j] == "" {
					count--
					j++
					continue
				}
				temp[i-j] = strings.Title(temp[i-j])
			}
			if comma {
				temp[i-1] = temp[i-1] + ","
			}
			i++
		case "(low,":
			count, comma, err := getNrepeats(text[i+1])
			errCheck(err)
			for j := 1; j <= count && i-j >= 0; j++ {
				if temp[i-j] == "" {
					count--
					j++
					continue
				}
				temp[i-j] = strings.ToLower(temp[i-j])
			}
			if comma {
				temp[i-1] = temp[i-1] + ","
			}
			i++
		case "(up,":
			count, comma, err := getNrepeats(text[i+1])
			errCheck(err)
			for j := 1; j <= count && i-j >= 0; j++ {
				if temp[i-j] == "" {
					count--
					j++
					continue
				}
				temp[i-j] = strings.ToUpper(temp[i-j])
			}
			if comma {
				temp[i-1] = temp[i-1] + ","
			}
			i++
		case "(bin)":
			n, err := strconv.ParseInt(temp[i-1], 2, 64)
			errCheck(err)
			temp[i-1] = strconv.Itoa(int(n))
		case "(hex)":
			n, err := strconv.ParseInt(temp[i-1], 16, 64)
			errCheck(err)
			temp[i-1] = strconv.Itoa(int(n))
		default:
			temp[i] = s
		}
	}
	for _, s := range temp {
		if s == "" {
			continue
		}
		correctedArray = append(correctedArray, s)
	}
	for _, s := range correctedArray[:len(correctedArray)-1] {
		result += s + " "
	}
	result += correctedArray[len(correctedArray)-1]

	return result
}

func getNrepeats(s string) (int, bool, error) {
	tail := s
	commacheck := false
	if tail[len(tail)-1] == ',' {
		commacheck = true
	}
	n := len(tail) - 1
	for ; n >= 0; n-- {
		if '0' <= tail[n] && tail[n] <= '9' {
			break
		}
	}
	count, err := strconv.Atoi(tail[:n+1])
	return count, commacheck, err
}

func articleChecker(last, curr string) string {
	firstChar := curr[0]
	newSentence := false
	if last[len(last)-1] == '.' {
		newSentence = true
	}
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, c := range vowels {
		if firstChar == c {
			if newSentence {
				return "An"
			} else {
				return "an"
			}
		}
	}
	if newSentence {
		return "A"
	} else {
		return "a"
	}
}
