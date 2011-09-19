package word

import (
	"testing"
)

var enableFilename = "../static/enable.txt"

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
	e := Factory(enableFilename)
	for _, word := range validWords {
		if !e.WordIsValid(word) {
			t.Errorf("%s was supposed to be in the dictionary\n", word)
		}
	}
	for _, word := range invalidWords {
		if e.WordIsValid(word) {
			t.Errorf("%s was NOT supposed to be in the dictionary\n", word)
		}
	}
}

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factory(enableFilename)
	}
}

func BenchmarkSearchLast(b *testing.B) {
	b.StopTimer()
	e := Factory(enableFilename)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e.WordIsValid("zyzzyvas")
	}
}

func BenchmarkSearchFail(b *testing.B) {
	b.StopTimer()
	e := Factory(enableFilename)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e.WordIsValid("yourmommaz")
	}
}
