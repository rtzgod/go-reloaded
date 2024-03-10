package internal

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func EditKeywords(input string) string {
	multipleCapPattern := regexp.MustCompile(`(\(\s*(cap|CAP)\s*,\s*-*[0-9]+\))`)
	multipleUpPattern := regexp.MustCompile(`(\(\s*(up|UP)\s*,\s*-*[0-9]+\))`)
	multipleLowPattern := regexp.MustCompile(`(\(\s*(low|LOW)\s*,\s*-*[0-9]+\))`)
	articlePattern := regexp.MustCompile(`(\s|^)(a|A|an|An)\s+[a-zA-Z]{3,}`)
	allInPattern := regexp.MustCompile(`\(hex\)|\(bin\)|\(up\)|\(low\)|\(cap\)|\(cap,\s*-*[0-9]+\)|\(up,\s*-*[0-9]+\)|\(low,\s*-*[0-9]+\)`)

	input = multipleCluFix(multipleCapPattern, input, "cap")
	input = multipleCluFix(multipleUpPattern, input, "up")
	input = multipleCluFix(multipleLowPattern, input, "low")

	text := strings.Fields(input)

	for i := 0; i < len(text); i++ {
		s := text[i]
		if len(s) >= 4 && s[:4] == "(up)" && i != 0 {
			text[i-1] = strings.ToUpper(text[i-1]) + s[4:]
			text = slices.Concat(text[:i], text[i+1:])
			i--
		} else if len(s) >= 5 && s[:5] == "(low)" && i != 0 {
			text[i-1] = strings.ToLower(text[i-1]) + s[5:]
			text = slices.Concat(text[:i], text[i+1:])
			i--
		} else if len(s) >= 5 && s[:5] == "(cap)" && i != 0 {
			text[i-1] = Capitalize(text[i-1]) + s[5:]
			text = slices.Concat(text[:i], text[i+1:])
			i--
		} else if len(s) >= 5 && s[:5] == "(bin)" && i != 0 {
			num, correct := numProcessing(text[i-1], "bin")
			if correct {
				text[i-1] = num + s[5:]
			}
			text = slices.Concat(text[:i], text[i+1:])
			i--
		} else if len(s) >= 5 && s[:5] == "(hex)" && i != 0 {
			num, correct := numProcessing(text[i-1], "hex")
			if correct {
				text[i-1] = num + s[5:]
			}
			text = slices.Concat(text[:i], text[i+1:])
			i--
		} else if len(s) >= 4 && s[:4] == "(up," && i != 0 {
			count, err := strconv.Atoi(getCount(s[:strings.LastIndex(s, ")")+1]))
			errCheck(err)
			if count <= 0 {
				text = slices.Concat(text[:i], text[i+1:])
				i--
				continue
			}
			text, i = multipleCluProcessing(text, s, "up", count, i)
		} else if len(s) >= 5 && s[:5] == "(low," && i != 0 {
			count, err := strconv.Atoi(getCount(s[:strings.LastIndex(s, ")")+1]))
			errCheck(err)
			if count <= 0 {
				text = slices.Concat(text[:i], text[i+1:])
				i--
				continue
			}
			text, i = multipleCluProcessing(text, s, "low", count, i)
		} else if len(s) >= 5 && s[:5] == "(cap," && i != 0 {
			count, err := strconv.Atoi(getCount(s[:strings.LastIndex(s, ")")+1]))
			errCheck(err)
			if count <= 0 {
				text = slices.Concat(text[:i], text[i+1:])
				i--
				continue
			}
			text, i = multipleCluProcessing(text, s, "cap", count, i)
		}
	}

	input = strings.Join(text, " ")
	input = articleProcessing(articlePattern, input)
	input = clearTrash(allInPattern, input)
	return input
}

func multipleCluFix(pattern *regexp.Regexp, s, kind string) string {
	return pattern.ReplaceAllStringFunc(s, func(match string) string {
		count := getCount(match)
		switch kind {
		case "up":
			return fmt.Sprintf("(up,%s)", count)
		case "low":
			return fmt.Sprintf("(low,%s)", count)
		case "cap":
			return fmt.Sprintf("(cap,%s)", count)
		default:
			return ""
		}
	})
}

func multipleCluProcessing(text []string, s, kind string, count, i int) ([]string, int) {
	if count > len(text[:i]) {
		count = len(text[:i])
	}
	for j := i - 1; count > 0; count-- {
		text[j] = cluFunc(text[j], kind)
		j--
	}
	text[i-1] = text[i-1] + s[strings.LastIndex(s, ")")+1:]
	text = slices.Concat(text[:i], text[i+1:])
	i--
	return text, i
}

func cluFunc(s, kind string) string {
	switch kind {
	case "up":
		return strings.ToUpper(s)
	case "low":
		return strings.ToLower(s)
	case "cap":
		return Capitalize(s)
	default:
		return ""
	}
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

func numProcessing(num, kind string) (string, bool) {
	switch kind {
	case "bin":
		num, err := strconv.ParseInt(num, 2, 64)
		if err != nil {
			return "", false
		}
		return strconv.Itoa(int(num)), true
	case "hex":
		num, err := strconv.ParseInt(num, 16, 64)
		if err != nil {
			return "", false
		}
		return strconv.Itoa(int(num)), true
	default:
		return "", false
	}
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

func charInSlice(char byte, arr []byte) bool {
	for _, element := range arr {
		if char == element {
			return true
		}
	}
	return false
}

func clearTrash(pattern *regexp.Regexp, s string) string {
	return pattern.ReplaceAllString(s, "")
}