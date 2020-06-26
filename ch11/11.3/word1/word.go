package word

func IsPalindrome(s string) bool {
	r := []rune(s)
	for i := range r {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}
