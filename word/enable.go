package word

import (
	"log"
	"os"
)

// ***************************************************************************

type WordList struct {
	words []string
}

func (p *WordList) Append(word string) {
	p.words = append(p.words, word)
}

func (p *WordList) WordInList(word string) bool {
	for _, w := range p.words {
		if w == word {
			return true
		}
	}
	return false
}

// ***************************************************************************

type Enable struct {
	mapWords map[int]WordList
	filename string
}

// singleton so more than one AppEngine request doesn't have to reload the file
// this should be in the AppEngine specific area, not here, because
// singletons are evil.  TODO: remove singleton
var singleton *Enable

func Factory(filename string) *Enable {
	if singleton != nil {
		return singleton
	}
	singleton = new(Enable)
	singleton.filename = filename
	singleton.loadDictionary()
	return singleton
}

// find at what seek point all the letters start
func (p *Enable) loadDictionary() {
	p.mapWords = make(map[int]WordList, 26)
	stream, err := os.OpenFile(p.filename, os.O_RDONLY, 0)
	defer stream.Close()
	if err != nil {
		log.Fatalf("failed to load dictionary %q", err)
	}
	pos := 0
	start := true // first byte read is an index
	bufferSize := 4096
	buffer := make([]byte, bufferSize)
	word := make([]byte, 0)
	for {
		cbuffer, err := stream.Read(buffer)
		if cbuffer == 0 || err == os.EOF {
			break
		} else if err != nil {
			log.Fatalf("failed to read from dictionary %q", err)
			return
		}
		for i := 0; i < cbuffer; i++ {
			if buffer[i] == '\n' || buffer[i] == '\r' || buffer[i] == 0 {
				// next is a start
				start = true
				continue
			} else if start && len(word) > 0 {
				start = false
				p.AddWord(string(word))
				word = make([]byte, 0)
			}
			// append any character that fell through
			word = append(word, buffer[i])
		}
		pos += cbuffer
	}
	p.AddWord(string(word))
}

func (p *Enable) AddWord(word string) {
	unicodeWord := []int(string(word)) // use []int instead of []byte to make unicode
	wl := p.mapWords[unicodeWord[0]]
	wl.Append(word)
	p.mapWords[unicodeWord[0]] = wl
}


// is this word in the Enable dictionary?
func (p *Enable) WordIsValid(query string) bool {
	if len(query) == 0 {
		return false
	}
	unicodeQuery := []int(query)
	thisChar := unicodeQuery[0]
	if p.mapWords == nil {
		log.Fatalf("dicationary not loaded, enable needs Init() %q", p)
	}
	wl := p.mapWords[thisChar]
	return wl.WordInList(query)
}
