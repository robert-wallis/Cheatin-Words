package word

import (
	"log"
	"os"
)

type Range struct {
	Start, End int
}

type Enable struct {
	m        map[int]Range
	filename string
}

var singleton *Enable

func (p Enable) Init(filename string) {
	if singleton != nil {
		return
	}
	p.filename = filename
	p.loadDictionary()
	singleton = &p
	return
}

// find at what seek point all the letters start
func (p Enable) loadDictionary() {
	p.m = make(map[int]Range, 26)
	stream, err := os.OpenFile(p.filename, os.O_RDONLY, 0)
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
					p.m[unicodeWord[0]] = Range{startIndex, 0}
					// modify previous range with end point
					p.updatePreviousEnd(lastChar, startIndex-1)
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
	p.updatePreviousEnd(lastChar, pos)
	stream.Close()
	log.Printf("%d word positions found\n", len(p.m))
}

func (p Enable) updatePreviousEnd(i, end int) {
	r, ok := p.m[i]
	if !ok {
		return
	}
	r.End = end
	p.m[i] = r
	log.Printf("%s = (%d, %d)\n", string(i), r.Start, r.End)
}

func StringInEnable(s string) bool {
	// TODO: os.OpenFile(filePath, os.O_RDONLY, 0)
	// stream.Seek(letterSeek[s[0]].Start)
	// reader.ReadString('\n')
	return false
}
