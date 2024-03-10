package internal

func CorrectAll(input string) string {
	input = CorrectPunctuation(input)
	input = EditKeywords(input)
	input = CorrectPunctuation(input)
	return input
}
