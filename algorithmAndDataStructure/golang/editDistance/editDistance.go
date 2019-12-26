package editDistance

func min(x int, others ...int) (minValue int) {
	minValue = x
	for _, other := range others {
		if minValue > other {
			minValue = other
		}
	}
	return
}

func minDistance(word1 string, word2 string) int {
	word1Len, word2Len := len(word1), len(word2)
	if word1Len == 0 {
		return word2Len
	}
	if word2Len == 0 {
		return word1Len
	}
	prevDp, curDp := make([]int, word1Len+1, word1Len+1), make([]int, word1Len+1, word1Len+1)

	for i := 0; i <= word1Len; i++ {
		prevDp[i] = i
	}

	for i := 1; i <= word2Len; i++ {
		curDp[0] = i
		for j := 1; j <= word1Len; j++ {
			if word1[j-1] == word2[i-1] {
				curDp[j] = prevDp[j-1]
			} else {
				curDp[j] = min(prevDp[j-1], prevDp[j], curDp[j-1]) + 1
			}
		}
		_ = copy(prevDp, curDp)
	}
	return prevDp[word1Len]
}
