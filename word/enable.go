package word

import (
	"log"
	"os"
)

type Range struct {
	Start, End int64
}

type Enable struct {
	m        map[int]Range
	filename string
}

var singleton *Enable

func (p *Enable) Init(filename string) {
	if singleton != nil {
		return
	}
	p.filename = filename
	p.loadDictionary()
	singleton = p
	return
}

// find at what seek point all the letters start
func (p *Enable) loadDictionary() {
	p.m = make(map[int]Range, 26)
	stream, err := os.OpenFile(p.filename, os.O_RDONLY, 0)
	defer stream.Close()
	if err != nil {
		log.Fatalf("failed to load dictionary %q", err)
	}
	pos := 0
	var lastChar int = 0
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
				// check for save
				unicodeWord := []int(string(word)) // use []int instead of []byte to make unicode
				if lastChar != unicodeWord[0] {
					// new char, time to save from the buffer's start + i - the beginning of this word in bytes
					startIndex := pos + i - len(word)
					// save
					p.m[unicodeWord[0]] = Range{int64(startIndex), 0}
					// modify previous range with end point
					p.updatePreviousEnd(lastChar, int64(startIndex-1))
					// save to compare with the next word
					lastChar = unicodeWord[0]
				}
				// if it was or wasn't this word, doesn't matter, new word
				word = make([]byte, 0)
			}
			// append any character that fell through
			word = append(word, buffer[i])
		}
		pos += cbuffer
	}
	p.updatePreviousEnd(lastChar, int64(pos))
	log.Printf("%d word positions found\n", len(p.m))
}

func (p *Enable) updatePreviousEnd(i int, end int64) {
	r, ok := p.m[i]
	if !ok {
		return
	}
	r.End = end
	p.m[i] = r
	log.Printf("%s = (%d, %d)\n", string(i), r.Start, r.End)
}

// is this word in the Enable dictionary?
func (p *Enable) WordIsValid(query string) bool {
	if len(query) == 0 {
		return false
	}
	unicodeQuery := []int(query)
	thisChar := unicodeQuery[0]
	r, ok := p.m[thisChar]
	if !ok {
		// word not in the index, therefore also not in dictionary
		return false
	}
	stream, err := os.OpenFile(p.filename, os.O_RDONLY, 0)
	defer stream.Close()
	if err != nil {
		log.Fatalf("failed to load dictionary %q", err)
		return false
	}
	pos, err := stream.Seek(r.Start, 0)
	if err != nil {
		log.Fatalf("couldn't seek to %d in index for %s %q", r.Start, string(thisChar), err)
		return false
	}
	start := true
	bufferSize := 4096
	buffer := make([]byte, bufferSize)
	word := make([]byte, 0)
	for {
		cbuffer, err := stream.Read(buffer)
		if cbuffer == 0 || err == os.EOF {
			break
		} else if err != nil {
			log.Fatalf("failed to read from dictionary %q", err)
			return false
		}
		for i := 0; i < cbuffer; i++ {
			if buffer[i] == '\n' || buffer[i] == '\r' || buffer[i] == 0 {
				// next is a start
				start = true
				continue
			} else if start && len(word) > 0 {
				start = false
				// check for save
				if string(word) == query {
					return true
				}
				// it wasn't this word, new word
				word = make([]byte, 0)
			}
			// append any character that fell through
			word = append(word, buffer[i])
		}
		pos += int64(cbuffer)
	}
	// check the very last word at the EOF
	if string(word) == query {
		return true
	}
	return false
}
