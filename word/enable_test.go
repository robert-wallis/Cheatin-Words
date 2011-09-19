package word

import (
	"fmt"
	"testing"
)

var validWords = []string{
	// regular words
	"abacus",
	"ball",
	"zygote",
	"a",        // first word ever
	"zyzzyvas", // last word ever
}

var invalidWords = []string{
	"yourmomma", // fake word
	"aaaa",      // string of letters
	"かんじ",       // non-ascii letters
	"汉字",        // non-ascii letters
}

func TestLoad(t *testing.T) {
	e := Factory("../static/enable.txt")
	for _, word := range validWords {
		fmt.Printf("%s ", word)
		if !e.WordIsValid(word) {
			t.Errorf("%s was supposed to be in the dictionary\n", word)
		}
		fmt.Printf("\n")
	}
	for _, word := range invalidWords {
		fmt.Printf("%s ", word)
		if e.WordIsValid(word) {
			t.Errorf("%s was NOT supposed to be in the dictionary\n", word)
		}
		fmt.Printf("\n")
	}
}
