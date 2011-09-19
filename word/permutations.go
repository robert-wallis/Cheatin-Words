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

// adopted from http://docs.python.org/library/itertools.html
func StringPermutationsSub(s string, r int) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		pool := []int(s) // get unicode
		n := len(pool)
		if r > n {
			return
		}
		indices := make([]int, n)
		for i := 0; i < n; i++ {
			indices[i] = i
		}
		cycles := []int{}
		for i := n; i > n-r; i-- {
			cycles = append(cycles, i)
		}
		endOfCycles := make([]int, len(cycles))
		copy(endOfCycles, cycles)
		// send the first iteration
		sout := make([]int, r)
		for i, v := range indices[:r] {
			sout[i] = pool[v]
		}
		out <- string(sout)
		for n > 0 {
			reversedRangeR := []int{}
			for i := r - 1; i >= 0; i-- {
				reversedRangeR = append(reversedRangeR, i)
			}
			for _, i := range reversedRangeR {
				cycles[i] = cycles[i] - 1
				if cycles[i] == 0 {
					// time for a new cycle
					newindices := make([]int, i)
					for k := 0; k < i; k++ {
						newindices[k] = indices[k]
					}
					for k := i + 1; k < len(indices); k++ {
						newindices = append(newindices, indices[k])
					}
					newindices = append(newindices, indices[i])
					indices = newindices
					cycles[i] = n - i
				} else {
					j := cycles[i]
					indices[i], indices[len(indices)-j] = indices[len(indices)-j], indices[i]
					// send iteration
					sout := make([]int, r)
					for k, v := range indices[:r] {
						sout[k] = pool[v]
					}
					out <- string(sout)
					break
				}
			}
			// have we come all the way around to the first permutation?
			same := true
			for i, v := range cycles {
				if endOfCycles[i] != v {
					same = false
					break
				}
			}
			if same {
				// yes this is the same as the first
				return
			}
		}
		// end
		return
	}()
	return out
}

