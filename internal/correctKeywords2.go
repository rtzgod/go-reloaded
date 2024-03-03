package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func EditKeywords2(input string) string {
	capPattern := regexp.MustCompile(`[!-~]+\s*\(cap\)`)
	lowPattern := regexp.MustCompile(`[!-~]+\s*\(low\)`)
	upPattern := regexp.MustCompile(`[!-~]+\s*\(up\)`)
	multipleCapPattern := regexp.MustCompile(`\(cap,\s*-*[0-9]+\)`)
	multipleLowPattern := regexp.MustCompile(`\(low,\s*-*[0-9]+\)`)
	multipleUpPattern := regexp.MustCompile(`\(up,\s*-*[0-9]+\)`)
	binPattern := regexp.MustCompile(`(\s|^)[01]+\s*\(bin\)`)
	hexPattern := regexp.MustCompile(`(\s|^)[0-9A-F]+\s*\(hex\)`)
	articlePattern := regexp.MustCompile(`(\s|^)(a|A|an|An)\s+[a-zA-Z]{3,}`)
	allInPattern := regexp.MustCompile(`\(hex\)|\(bin\)|\(up\)|\(low\)|\(cap\)|\(cap,\s*-*[0-9]+\)|\(up,\s*-*[0-9]+\)|\(low,\s*-*[0-9]+\)`)

	result := cluProcessing(upPattern, "up", input)
	result = cluProcessing(lowPattern, "low", result)
	result = cluProcessing(capPattern, "cap", result)
	result = multipleCluProcessing(multipleUpPattern, "up", result)
	result = multipleCluProcessing(multipleLowPattern, "low", result)
	result = multipleCluProcessing(multipleCapPattern, "cap", result)
	result = baseProcessing(binPattern, "bin", result)
	result = baseProcessing(hexPattern, "hex", result)
	result = articleProcessing(articlePattern, result)
	result = clearTrash(allInPattern, result)
	return result
}

func cluProcessing(pattern *regexp.Regexp, kind, s string) string {
	return CorrectPunctuation(pattern.ReplaceAllStringFunc(s, func(match string) string {
		word := extractValue(match)
		word = wordProcessing(word, kind)
		return word
	}))
}

func multipleCluProcessing(pattern *regexp.Regexp, kind, s string) string {
	idx := pattern.FindStringIndex(s)
	for idx != nil {
		count, _ := strconv.Atoi(getCount(s[idx[0]:idx[1]]))
		if count <= 0 {
			s = s[:idx[0]] + s[idx[1]:]
			s = CorrectPunctuation(s)
			idx = pattern.FindStringIndex(s)
			continue
		}
		new_s := strings.Fields(s[:idx[0]])
		if count > len(new_s) {
			count = len(new_s)
		}
		new_s = new_s[len(new_s)-count:]
		fmt.Println(new_s)
		old_s := strings.Join(new_s, " ")
		for i := 0; i < len(new_s); i++ {
			new_s[i] = wordProcessing(new_s[i], kind)
		}
		result := strings.Join(new_s, " ")
		s = s[:idx[0]] + s[idx[1]:]
		s = CorrectPunctuation(s)
		s = strings.Replace(s, old_s, result, 1)
		idx = pattern.FindStringIndex(s)
	}
	return s
}

func baseProcessing(pattern *regexp.Regexp, kind, s string) string {
	return CorrectPunctuation(pattern.ReplaceAllStringFunc(s, func(match string) string {
		num := extractValue(match)
		num = numProcessing(num, kind)
		if match[0] == ' ' {
			return " " + num
		}
		return num
	}))
}

func articleProcessing(pattern *regexp.Regexp, s string) string {
	return CorrectPunctuation(pattern.ReplaceAllStringFunc(s, func(match string) string {
		article := getArticle(match)
		word := getWordAfterArticle(match)

		vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

		var vowel bool = charInSlice(word[0], vowels)

		switch article {
		case "a", "A":
			if vowel {
				article = article + "n"
			}
		case "an", "An":
			if !vowel {
				article = string(article[0])
			}
		}
		return " " + article + " " + word
	}))
}

func charInSlice(char byte, arr []byte) bool {
	for _, element := range arr {
		if char == element {
			return true
		}
	}
	return false
}

func extractValue(match string) string {
	var value string
	if match[0] == ' ' {
		match = match[1:]
	}
	for _, c := range match {
		if c == ' ' || c == '(' {
			break
		}
		value += string(c)
	}
	return value
}

func wordProcessing(word, kind string) string {
	switch kind {
	case "cap":
		return Capitalize(word)
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	default:
		return word
	}
}

func numProcessing(num, kind string) string {
	switch kind {
	case "bin":
		num, err := strconv.ParseInt(num, 2, 64)
		errCheck(err)
		return strconv.Itoa(int(num))
	case "hex":
		num, err := strconv.ParseInt(num, 16, 64)
		errCheck(err)
		return strconv.Itoa(int(num))
	default:
		return num
	}
}

func getArticle(match string) string {
	var article string
	if match[0] == ' ' {
		match = match[1:]
	}
	for _, c := range match {
		if c == ' ' {
			break
		}
		article += string(c)
	}
	return article
}

func getCount(match string) string {
	end := len(match) - 1
	start := end
	for ; start >= 0; start-- {
		if match[start] == ' ' || match[start] == ',' {
			break
		}
	}
	return match[start+1 : end]
}

func getWordAfterArticle(match string) string {
	var word string
	for i := len(match) - 1; i >= 0; i-- {
		if match[i] == ' ' {
			break
		}
		word = string(match[i]) + word
	}
	return word
}

func clearTrash(pattern *regexp.Regexp, s string) string {
	return pattern.ReplaceAllString(s, "")
}
