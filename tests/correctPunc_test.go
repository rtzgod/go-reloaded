package main

import (
	inter "go_reloaded/internal"
	"testing"
)

func TestCorrectPunctuation(t *testing.T) {
	var tests = []struct {
		id    string
		input string
		want  string
	}{
		{"!! test", "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"... test", "I was thinking ... You were right", "I was thinking... You were right"},
		{" 'one word' test", "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{" 'multiple words' test", "As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			ans := inter.CorrectPunctuation(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
