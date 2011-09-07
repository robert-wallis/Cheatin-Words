package word

import (
	"sort"
	"testing"
)

type permuteTest struct {
	in, out sort.Interface
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
		if []int(pt.in) != []int(pt.out) {
			t.Error("expected Permute() value %q, got %q", pt.in, pt.out)
		}
	}
}
