package word

import (
	"os"
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
	e := new(Enable)
	if err := e.Load(enableFilename); err != nil {
		t.Error("unable to load enable for filename %s: %s", enableFilename, err)
	}
	for _, word := range validWords {
		valid := true
		var err os.Error
		if valid, err = e.WordIsValid(word); !valid {
			t.Errorf("%s was supposed to be in the dictionary\n", word)
		}
		if err != nil {
			t.Errorf("error searching for %s: %s", word, err)
		}
	}
	for _, word := range invalidWords {
		valid := false
		var err os.Error
		if valid, err = e.WordIsValid(word); valid {
			t.Errorf("%s was NOT supposed to be in the dictionary\n", word)
		}
		if err != nil {
			t.Errorf("error searching for %s: %s", word, err)
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
