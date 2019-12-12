package twoSum

import (
	"sort"
)

// By is the type of a "less" function that defines the ordering of its Number arguments
type By func(n1, n2 *Number) bool

// Number defines the properties of a number object
type Number struct {
	Val int
	Pos int
}

// Sort is a method on the function type, By, that sorts the argument slice according to the function
func (by By) Sort(numbers []*Number) {
	ns := &numberSorter{
		nums: numbers,
		by:   by,
	}
	sort.Sort(ns)
}

// numberSorter joins a By function and a slice of Numbers to be sorted
type numberSorter struct {
	nums []*Number
	by   func(n1, n2 *Number) bool
}

// Len is part of sort.Interface
func (s *numberSorter) Len() int {
	return len(s.nums)
}

// Swap is part of sort.Interface
func (s *numberSorter) Swap(i, j int) {
	s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *numberSorter) Less(i, j int) bool {
	return s.by(s.nums[i], s.nums[j])
}

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	ns := make([]*Number, 0, numsLen)
	for i := 0; i < numsLen; i++ {
		ns = append(ns, &Number{
			Val: nums[i],
			Pos: i,
		})
	}
	value := func(n1, n2 *Number) bool {
		return n1.Val < n2.Val
	}
	By(value).Sort(ns)

	res := make([]int, 0, 2)
	for i, j := 0, numsLen-1; i < j; {
		sum := ns[i].Val + ns[j].Val
		if sum == target {
			res = append(res, ns[i].Pos, ns[j].Pos)
			break
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return res
}
