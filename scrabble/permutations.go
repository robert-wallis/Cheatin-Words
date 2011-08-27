package scrabble

/**
 * get a channel of all permutations of n objects
 * ex. n = 3 then [0 1 2] [0 2 1] [1 0 2] [1 2 0] [2 0 1] [2 1 0] nil
 * use the yeilded slice as indexes of your object's slice
 * channel yeilds nil when done
 */
func Permutations(n int) chan []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i
	}
	yeild := make(chan []int)
	go func() {
		// must make a copy because the next permutation overwrites the same memory
		// and the channel won't block until after it is full
		result := make([]int, n)
		copy(result, list)
		yeild <- result
		for nextPermutation(list) {
			yeild <- list
		}
		close(yeild)
	}()
	return yeild
}

// changes the order of the sequence in memory to the next permutation
func nextPermutation(seq []int) bool {
	for j := len(seq) - 1; j > 0; j-- {
		if v := seq[j-1]; v < seq[j] {
			m := len(seq) - 1
			for v > seq[m] {
				m--
			}
			seq[j-1], seq[m] = seq[m], seq[j-1]
			reverse(seq[j:])
			return true
		}
	}
	return false
}

func reverse(seq []int) {
	for i, j := 0, len(seq)-1; i < j; i, j = i+1, j-1 {
		seq[i], seq[j] = seq[j], seq[i]
	}
}

/**
 * returns a channel of all possible permutations for that string
 */
func StringPermutations(s string) chan string {
	out := make(chan string)
	go func() {
		perms := Permutations(len(s))
		for p := range perms {
			out <- rearrangeString(s, p)
		}
		close(out)
	}()
	return out
}

// re-arrange string according to the index order in arrangement
func rearrangeString(s string, arrangement []int) string {
	cp := make([]byte, len(arrangement))
	for i, c := range arrangement {
		cp[i] = s[c]
	}
	return string(cp)
}
