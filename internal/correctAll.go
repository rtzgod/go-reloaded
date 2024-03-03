package internal

func CorrectAll(input string) string {
	input = CorrectPunctuation(input)
	input = EditKeywords2(input)
	input = CorrectPunctuation(input)
	return input
}
