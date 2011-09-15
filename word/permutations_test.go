package word

import (
	"testing"
)

type permuteTest struct {
	in, out IntSlice
	ret     bool	// are there more permutations after this?
}

var permuteTests = []*permuteTest{
	// regular tests
	&permuteTest{IntSlice{1, 2}, IntSlice{2, 1}, true},
	&permuteTest{IntSlice{2, 1}, IntSlice{2, 1}, false},
	// basic ab ba test
	&permuteTest{IntSliceFromString("ab"), IntSliceFromString("ba"), true},
	&permuteTest{IntSliceFromString("ba"), IntSliceFromString("ba"), false},
	// advanced aab test to make sure it doesn't consider each 'a' as seperate
	// but that 'a₁a₂b₃' == 'a₂a₁b₃' so there are pragmatically less than N! permutations
	&permuteTest{IntSliceFromString("aab"), IntSliceFromString("aba"), true},
	&permuteTest{IntSliceFromString("aba"), IntSliceFromString("baa"), true},
	&permuteTest{IntSliceFromString("baa"), IntSliceFromString("baa"), false},
	// just double checking that last one with a double second character instead
	&permuteTest{IntSliceFromString("bba"), IntSliceFromString("bba"), false},
	// unicode tests
	&permuteTest{IntSliceFromString("一二"), IntSliceFromString("二一"), true},
	&permuteTest{IntSliceFromString("二一"), IntSliceFromString("二一"), false},
}

func TestPermute(t *testing.T) {
	for _, pt := range permuteTests {
		ret := Permute(pt.in)
		if ret != pt.ret {
			t.Errorf("expected Permute() return %q, got %q", pt.ret, ret)
		}
		if len(pt.in) != len(pt.out) {
			t.Errorf("expeted Permute to have %d length, but had %d", len(pt.in), len(pt.out))
		}
		for i, v := range pt.in {
			if pt.out[i] != v {
				t.Errorf("expected Permute() index %d to be %q, but got %q", i, v, pt.out[i])
			}
		}
	}
}
