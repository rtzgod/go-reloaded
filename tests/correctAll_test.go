package main

import (
	inter "go_reloaded/internal"
	"testing"
)

func TestCorrectAll(t *testing.T) {
	var tests = []struct {
		id    string
		input string
		want  string
	}{
		{"1", "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.", "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."},
		{"2", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
		{"3", "There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
		{"4", "Punctuation tests are ... kinda boring ,don't you think !?", "Punctuation tests are... kinda boring, don't you think!?"},
		{"5", "Ready, set, go (up) ?? ? DOESN'T WORKS", "Ready, set, GO??? DOESN'T WORKS"},
		{"6", "a a a a a apple", "a a a a an apple"},
		{"7", "' cat '  (cap)", "'Cat'"},
		{"8", "fsdaf fsfaa fds A APLLE", "fsdaf fsfaa fds An APLLE"},
		{"9", "fsdo (cap, -3)", "fsdo"},
		{"10", "//dfsuhv (cap, 0)", "//dfsuhv"},
		{"11", "hdfuigos h45789 (cap)", "hdfuigos H45789"},
		{"12", "111 (bin)", "7"},
		{"13", "ghfug\ngfuh\ngfdi\nghfd\n(up,6)", "GHFUG GFUH GFDI GHFD"},
		{"14", "This is word09 (hex)", "This is word09"},
		{"15", "This is word01 (bin)", "This is word01"},
		{"16", "This is another several words (up, 4) (low,2) (cap, 3)", "This IS Another Several Words"},
		{"17", "", ""},
		{"18", "(hex)", ""},
		{"19", "(bin)", ""},
		{"20", "(up)", ""},
		{"21", "(low)", ""},
		{"22", "(cap)", ""},
		{"23", "1E (hex) files were added", "30 files were added"},
		{"24", "It has been 10 (bin) years", "It has been 2 years"},
		{"25", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"26", "Ready, set, go\n(up) !", "Ready, set, GO!"},
		{"27", "I should stop SHOUTING (low)", "I should stop shouting"},
		{"28", "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"29", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"30", "This is so exciting (up, -2)", "This is so exciting"},
		{"31", "This is So ExciTIng (low, 2)", "This is so exciting"},
		{"32", "This is so exciting (cap, 2)", "This is So Exciting"},
		{"33", "This is so exciting (cap, 10)", "This Is So Exciting"},
		{"34", "it was the best of times, it was the worst of times, it was the age of  wisdom,  it was the age of  foolishness (up, 15)", "it was the best of times, it was the WORST OF TIMES, IT WAS THE AGE OF WISDOM, IT WAS THE AGE OF FOOLISHNESS"},
		{"35", "I should stop SHOUTING (low, 999)", "i should stop shouting"},
		{"36", "it was the best of times, it was the WORST OF TIMES, IT WAS  THE AGE OF WISDOM, IT WAS THE  AGE OF FOOLISHNESS (low, 15)", "it was the best of times, it was the worst of times, it was the age of wisdom, it was the age of foolishness"},
		{"37", "it was the best of times, it was the worst of times, it was the age of  wisdom, it was the age of  foolishness (cap, 15)", "it was the best of times, it was the Worst Of Times, It Was The Age Of Wisdom, It Was The Age Of Foolishness"},
		{"38", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"39", "fsdaf fsfaa fds A APLLE", "fsdaf fsfaa fds An APLLE"},
		{"40", "fdsfsdf (cap, 4", "fdsfsdf (cap, 4"},
		{"41", "valera(cap)", "Valera"},
		{"42", "(cap)(up)(low)", ""},
		{"43", "valera(cap)(low)(up)", "VALERA"},
		{"44", "valera(cap, 1)(low, 3)(up, 4)", "VALERA"},
		{"45", "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?", "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"46", "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure", "I have to pack 5 outfits. Packed 26 just to be sure"},
		{"47", "Don not be sad ,because sad backwards is das . And das not good", "Don not be sad, because sad backwards is das. And das not good"},
		{"48", "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '", "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			ans := inter.CorrectAll(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
