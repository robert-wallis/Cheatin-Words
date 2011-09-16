package word

import (
	"testing"
)

func TestLoad(t *testing.T) {
	e := new(Enable)
	e.Init("../static/enable.txt")
}
