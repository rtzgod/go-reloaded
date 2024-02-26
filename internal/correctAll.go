package internal

func CorrectAll(input string) string {
	result := EditKeywords(input)
	result = CorrectPunctuation(result)
	return result
}
