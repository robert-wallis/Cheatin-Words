package word

import (
	"fmt"
	"os"
)

// ***************************************************************************

type Enable struct {
	mapWords map[string]bool
	filename string
}

// find at what seek point all the letters start
func (p *Enable) Load(filename string) os.Error {
	p.filename = filename
	p.mapWords = make(map[string]bool, 26)
	stream, err := os.OpenFile(p.filename, os.O_RDONLY, 0)
	defer stream.Close()
	if err != nil {
		return os.NewError(fmt.Sprintf("failed to load dictionary %q", err))
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
			return os.NewError(fmt.Sprintf("failed to read from dictionary %q", err))
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
	return nil
}

func (p *Enable) AddWord(word string) {
	p.mapWords[word] = true
}


// is this word in the Enable dictionary?
func (p *Enable) WordIsValid(query string) bool {
	if len(query) == 0 {
		return false
	}
	if _, inMap := p.mapWords[query]; !inMap {
		return false
	}
	return true
}
