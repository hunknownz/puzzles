package longestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	const (
		maxChars = 256
	)
	var alphabet [maxChars]int
	for i := 0; i < len(alphabet); i++ {
		alphabet[i] = -1
	}

	lastPos, maxLen := -1, 0

	for i, char := range s {
		value := int(char)
		if alphabet[value] != -1 && lastPos < alphabet[value] {
			lastPos = alphabet[value]
		}

		alphabet[value] = i

		if maxLen < i-lastPos {
			maxLen = i - lastPos
		}
	}
	return maxLen
}
