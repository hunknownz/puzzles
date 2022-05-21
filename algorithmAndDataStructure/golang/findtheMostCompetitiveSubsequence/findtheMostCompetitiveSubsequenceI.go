package findtheMostCompetitiveSubsequence

func mostCompetitiveI(nums []int, k int) []int {
	numLen := len(nums)

	ms := NewMonotoneStack(k)

	for i, num := range nums {
		remainLen := numLen - i - 1
		ms.Push(num, remainLen)
	}

	return ms.Stack()
}

type MonotoneStack struct {
	size  int
	stack []int
	p     int
}

func (ms *MonotoneStack) Push(val, remainLen int) {
	for ms.p != -1 && remainLen+ms.p+1 >= ms.size && val < ms.stack[ms.p] {
		ms.p--
	}
	if ms.p == ms.size-1 {
		return
	}
	ms.p++
	ms.stack[ms.p] = val
}

func (ms *MonotoneStack) Stack() (stack []int) {
	stack = make([]int, ms.size)
	copy(stack, ms.stack)
	return
}

func NewMonotoneStack(size int) (ms *MonotoneStack) {
	ms = new(MonotoneStack)
	ms.p = -1
	ms.size = size
	ms.stack = make([]int, size, size)

	return
}
