package validParentheses

var pMap map[rune]rune = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func isValid(s string) bool {
	sLen := len(s)
	stack := make([]rune, sLen, sLen)
	p := -1

	for _, par := range s {
		if _, ok := pMap[par]; ok {
			p++
			stack[p] = par
			continue
		}

		if p != -1 && pMap[stack[p]] == par {
			p--
		} else {
			return false
		}
	}

	return p == -1
}
