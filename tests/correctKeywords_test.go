package main

import (
	inter "go_reloaded/internal"
	"testing"
)

func TestEditKeywords(t *testing.T) {
	var tests = []struct {
		id    string
		input string
		want  string
	}{
		{"(hex) test", "1E (hex) files were added", "30 files were added"},
		{"(bin) test", "It has been 10 (bin) years", "It has been 2 years"},
		{"(up) test", "Ready, set, go (up) !", "Ready, set, GO !"},
		{"(low) test", "I should stop SHOUTING (low)", "I should stop shouting"},
		{"(cap) test", "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"(up, number) test", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"(low, number) test", "This is So ExciTIng (low, 2)", "This is so exciting"},
		{"(cap, number) test", "This is so exciting (cap, 2)", "This is So Exciting"},
		{"(cap, number) overflow test", "This is so exciting (cap, 10)", "This Is So Exciting"},
		{"(up, number) big number test", "it was the best of times, it was the worst of times, it was the age of  wisdom,  it was the age of  foolishness (up, 15)", "it was the best of times, it was the WORST OF TIMES, IT WAS THE AGE OF WISDOM, IT WAS THE AGE OF FOOLISHNESS"},
		{"(low, number) big number test", "it was the best of times, it was the WORST OF TIMES, IT WAS  THE AGE OF WISDOM, IT WAS THE  AGE OF FOOLISHNESS (low, 15)", "it was the best of times, it was the worst of times, it was the age of wisdom, it was the age of foolishness"},
		{"(cap, number) big number test", "it was the best of times, it was the worst of times, it was the age of  wisdom, it was the age of  foolishness (cap, 15)", "it was the best of times, it was the Worst Of Times, It Was The Age Of Wisdom, It Was The Age Of Foolishness"},
		{"a->an test", "There it was. A amazing rock!", "There it was. An amazing rock!"},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			ans := inter.EditKeywords(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
