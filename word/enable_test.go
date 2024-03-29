package word

import (
	"testing"
)

var enableFilename = "../enable.txt"

var validWords = []string{
	// regular words
	"abacus",
	"ball",
	"zygote",
	"aa",       // first word
	"zyzzyvas", // last word
}

var invalidWords = []string{
	"yourmomma", // fake word
	"aaaa",      // string of letters
	"かんじ",       // non-ascii letters
	"汉字",        // non-ascii letters
}

func TestLoad(t *testing.T) {
	e := new(Enable)
	if err := e.Load(enableFilename); err != nil {
		t.Error("unable to load enable for filename %s: %s", enableFilename, err)
	}
	for _, word := range validWords {
		if valid := e.WordIsValid(word); !valid {
			t.Errorf("%s was supposed to be in the dictionary\n", word)
		}
	}
	for _, word := range invalidWords {
		if valid := e.WordIsValid(word); valid {
			t.Errorf("%s was NOT supposed to be in the dictionary\n", word)
		}
	}
}

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e := new(Enable)
		e.Load(enableFilename)
	}
}

func BenchmarkSearchLast(b *testing.B) {
	b.StopTimer()
	e := new(Enable)
	e.Load(enableFilename)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e.WordIsValid("zyzzyvas")
	}
}

func BenchmarkSearchFail(b *testing.B) {
	b.StopTimer()
	e := new(Enable)
	e.Load(enableFilename)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e.WordIsValid("yourmommaz")
	}
}
