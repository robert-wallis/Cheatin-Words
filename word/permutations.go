package word

import (
	"sort"
)

// finds the next permutation and sorts it in data
// returns false when there are not permutations left
// based on: http://en.wikipedia.org/wiki/Permutation#Systematic_generation_of_all_permutations
// based on: https://github.com/cznic/mathutil/blob/master/permute.go
func Permute(data sort.Interface) bool {
	var i, j int
	// check if it's already sorted
	for i = data.Len() - 2; ; i-- {
		if i < 0 {
			return false
		}
		if data.Less(i, i+1) {
			break
		}
	}
	// find the next j to swap
	for j = data.Len() - 1; !data.Less(i, j); j-- {
	}
	// do this swap
	data.Swap(i, j)
	// carry swap
	for i, j := i+1, data.Len()-1; i < j; i++ {
		data.Swap(i, j)
		j--
	}
	return true
}

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
// from new sort, but not in google_appengine_go sort
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p IntSlice) Sort() { sort.Sort(p) }
func IntSliceFromString(s string) IntSlice {
	return []int(s)
}
func StringFromIntSlice(i IntSlice) string {
	return string([]int(i))
}
// ---------------------------------------------------------------------------

// returns a channel of all the possible permutations of a string
func StringPermutations(s string) chan string {
	out := make(chan string)
	go func() {
		// make a mutable string
		b := IntSliceFromString(s)
		// start the permutations at the beginning
		sort.Sort(b)
		for {
			out <- StringFromIntSlice(b)
			if ok := Permute(b); !ok {
				break
			}
		}
		close(out)
	}()
	return out
}
