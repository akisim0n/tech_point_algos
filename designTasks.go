package main

import "math"

type MinStack struct {
	nums  []int
	minEl int
}

func Constructor() MinStack {
	return MinStack{minEl: math.MaxInt}
}

func (s *MinStack) Push(val int) {

	s.nums = append(s.nums, val)

	if val < s.minEl {
		s.minEl = val
	}
}

func (s *MinStack) Pop() {
	removedValue := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]

	if removedValue == s.minEl && len(s.nums) > 0 {
		tmpMin := s.nums[0]
		for i := 1; i < len(s.nums); i++ {
			if s.nums[i] < tmpMin {
				tmpMin = s.nums[i]
			}
		}
		s.minEl = tmpMin
	} else if len(s.nums) == 0 {
		s.minEl = math.MaxInt
	}
}

func (s *MinStack) Top() int {
	return s.nums[len(s.nums)-1]
}

func (s *MinStack) GetMin() int {
	return s.minEl
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
