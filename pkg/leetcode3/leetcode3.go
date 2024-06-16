package leetcode3

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func Solve(s string) int {
	maxLen := 0
	visited := make(map[rune]int, len(s)-1)
	currentLen := 0
	indexStart := 0
	for index, char := range s {
		if matchedIndex, ok := visited[char]; ok {
			for indexStart <= matchedIndex {
				keyToRemove := s[indexStart]
				delete(visited, rune(keyToRemove))
				indexStart++
			}
			currentLen = index - (matchedIndex + 1)
		}
		currentLen++
		visited[char] = index
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
