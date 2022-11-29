package util

import "math"

// Returns maximun value of a list of integers. If no element is insert, it returns minimun value available
func Max(values ...int) int {
	if len(values) == 0 {
		return math.MinInt
	}
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

// Returns minimun value of a list of integers. If no element is insert, it returns maximun value available
func Min(values ...int) int {
	if len(values) == 0 {
		return math.MaxInt
	}
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}
