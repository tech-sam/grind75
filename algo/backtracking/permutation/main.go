package main

func main() {
	s := "abc"
	find_permutations(s, 0, len(s)-1)

}

func find_permutations(s string, left, right int) {

	for i := left; i <= right; i++ {

		s = swap(s, left, i)
		find_permutations(s, left+1, i)
		s = swap(s, left, i)
	}
}

func swap(s string, left, i int) string {
	// t := s[left]
	// s[i] = s[left]
	return ""
}
