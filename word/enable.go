package word

import (
	"fmt"
	"os"
	"sort"
	"io/ioutil"
	"bytes"
)

// ***************************************************************************

type Enable struct {
	words    [][]byte
	filename string
}

func (p *Enable) Load(filename string) os.Error {
	p.filename = filename

	data, err := ioutil.ReadFile(p.filename)
	if err != nil {
		return os.NewError(fmt.Sprintf("failed to load dictionary %q", err))
	}
	p.words = bytes.Split(data, []byte("\n"), -1)
	return nil
}

// is this word in the Enable dictionary?
func (p *Enable) WordIsValid(query string) bool {
	if len(query) == 0 {
		return false
	}
	bquery := []byte(query)
	i := sort.Search(len(p.words), func(i int) bool {
		return bytes.Compare(p.words[i], bquery) >= 0
	})
	return i < len(p.words) && bytes.Compare(p.words[i], bquery) == 0
}
