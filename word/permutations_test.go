package word

import (
	"sort"
	"testing"
)

type permuteTest struct {
	in, out sort.IntSlice
	ret     bool
}

var permuteTests = []*permuteTest{
	&permuteTest{sort.IntSlice{1, 2}, sort.IntSlice{2, 1}, true},
	&permuteTest{sort.IntSlice{2, 1}, sort.IntSlice{2, 1}, false},
}

func TestPermute(t *testing.T) {
	for _, pt := range permuteTests {
		ret := Permute(pt.in)
		if ret != pt.ret {
			t.Error("expected Permute() return %q, got %q", pt.ret, ret)
		}
		if len(pt.in) != len(pt.out) {
			t.Error("expeted Permute to have %d length, but had %d", len(pt.in), len(pt.out))
		}
		for i, v := range pt.in {
			if pt.out[i] != v {
				t.Error("expected Permute() index %d to be %q, but got %q", i, v, pt.out[i])
			}
		}
	}
}
