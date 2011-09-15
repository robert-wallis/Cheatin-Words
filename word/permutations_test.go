package word

import (
	"fmt"
	"sort"
	"testing"
)

type intPermuteTest struct {
	in, out sort.IntSlice
	ret     bool
}

var intPermuteTests = []*intPermuteTest{
	&intPermuteTest{sort.IntSlice{1, 2}, sort.IntSlice{2, 1}, true},
	&intPermuteTest{sort.IntSlice{2, 1}, sort.IntSlice{2, 1}, false},
}

func TestIntPermute(t *testing.T) {
	for _, pt := range intPermuteTests {
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

type stringPermuteTest struct {
	in, out ByteSlice
	ret     bool
}

var stringPermuteTests = []*stringPermuteTest{
	&stringPermuteTest{ByteSliceFromString("ab"), ByteSliceFromString("ba"), true},
	&stringPermuteTest{ByteSliceFromString("ba"), ByteSliceFromString("ba"), false},
	&stringPermuteTest{ByteSliceFromString("aab"), ByteSliceFromString("aba"), true},
	&stringPermuteTest{ByteSliceFromString("aba"), ByteSliceFromString("baa"), true},
	&stringPermuteTest{ByteSliceFromString("baa"), ByteSliceFromString("baa"), false},
	&stringPermuteTest{ByteSliceFromString("bba"), ByteSliceFromString("bba"), false},
}

func TestStringPermute(t *testing.T) {
	for i, pt := range stringPermuteTests {
		fmt.Println(i)
		fmt.Println("before", pt.in)
		ret := Permute(pt.in)
		fmt.Println("after", pt.in)
		if ret != pt.ret {
			fmt.Println("fail", ret, pt.ret)
			t.Errorf("expected Permute() return %q, got %q", pt.ret, ret)
		}
		if len(pt.in) != len(pt.out) {
			fmt.Println("fail", len(pt.in), len(pt.out))
			t.Errorf("expeted Permute to have %d length, but had %d", len(pt.in), len(pt.out))
		}
		for i, v := range pt.in {
			if pt.out[i] != v {
				fmt.Println("fail", pt.in, pt.out)
				t.Errorf("expected Permute() index %d to be %q, but got %q", i, v, pt.out[i])
			}
		}
	}
}
