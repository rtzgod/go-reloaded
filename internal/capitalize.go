package internal

func Capitalize(s string) string {
	var new_s string
	needToUp := true
	needToLow := false
	for _, c := range s {
		if !IsAlphaRune(c) {
			needToUp = true
			needToLow = false
		}
		if needToUp && (c >= 'a' && c <= 'z') {
			c = c - 32
			needToLow = true
			needToUp = false
			new_s += string(c)
			continue
		} else if (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			needToUp = false
		}
		if needToLow && (c >= 'A' && c <= 'Z') {
			c = c + 32
			new_s += string(c)
			continue
		}
		if !needToLow && !needToUp {
			needToLow = true
		}
		new_s += string(c)
	}
	return new_s
}
func IsAlphaRune(c rune) bool {
	return ((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9'))
}
