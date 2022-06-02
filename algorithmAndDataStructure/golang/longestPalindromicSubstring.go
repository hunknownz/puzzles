package longestPalindromicSubstring

func longestPalindrome(s string) string {
	l, r := manacher(s)
	return s[l : r+1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func manacher(s string) (l, r int) {
	stringLen := len(s)
	d := make([]int, stringLen, stringLen)
	maxOdd, maxEven := -1, -1
	var maxOddIndex, maxEvenIndex int

	for i, l, r := 0, 0, -1; i < stringLen; i++ {
		var k int
		if i > r {
			k = 1
		} else {
			k = min(d[l+r-i], r-i+1)
		}
		for 0 <= i-k && i+k < stringLen && s[i-k] == s[i+k] {
			k++
		}
		d[i] = k
		if k > maxOdd {
			maxOdd = k
			maxOddIndex = i
		}
		if i+k-1 > r {
			l = i - k + 1
			r = i + k - 1
		}
	}

	for i, l, r := 0, 0, -1; i < stringLen; i++ {
		var k int
		if i > r {
			k = 0
		} else {
			k = min(d[l+r-i+1], r-i+1)
		}
		for 0 <= i-k-1 && i+k < stringLen && s[i-k-1] == s[i+k] {
			k++
		}
		d[i] = k
		if k > maxEven {
			maxEven = k
			maxEvenIndex = i
		}
		if i+k-1 > r {
			l = i - k
			r = i + k - 1
		}
	}

	if maxOdd > maxEven {
		l = maxOddIndex - maxOdd + 1
		r = maxOddIndex + maxOdd - 1
	} else {
		l = maxEvenIndex - maxEven
		r = maxEvenIndex + maxEven - 1
	}
	return
}
