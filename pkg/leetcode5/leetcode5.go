package leetcode5

func isPalindrome(it string, itLen int) bool {
	if itLen == 1 {
		return true
	}
	var end int
	if itLen%2 == 0 {
		end = itLen / 2
	} else {
		end = (itLen - 1) / 2
	}
	for i := 0; i < end; i++ {
		altIndex := itLen - 1 - i
		if it[i] != it[altIndex] {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/longest-palindromic-substring/description/
func Solve(s string) string {
	sLength := len(s)
	if sLength == 0 {
		return ""
	}
	longestPalindrome := s[0:1]
	longestPalindromeLength := 0
	for i := 1; i < sLength; i++ {
		end := i + 1
		for start := 0; start < i; start++ {
			wordLength := end - start
			if wordLength < longestPalindromeLength {
				continue
			}
			word := s[start:end]
			if isPalindrome(word, wordLength) {
				if wordLength > longestPalindromeLength {
					longestPalindrome = word
					longestPalindromeLength = wordLength
				}
			}
		}
	}
	return longestPalindrome
}
