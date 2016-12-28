package main

// MaxPairwiseProduct finds the maximum pairwise product given a sequence of
// non-negative integers, that is, the largest integer that can be obtained by
// multiplying two different elements from the sequence.
func MaxPairwiseProduct(seq []int) int {
	tmp := make([]int, len(seq))
	copy(tmp, seq)
	index1 := maxElementIndex(tmp)
	max1 := tmp[index1]
	tmp = append(tmp[:index1], tmp[index1+1:]...)
	index2 := maxElementIndex(tmp)
	return max1 * tmp[index2]
}

func maxElementIndex(seq []int) (max int) {
	for i, v := range seq {
		if v > seq[max] {
			max = i
		}
	}
	return
}
