package util

// Returns maximun value of a list of integers
func Max(values ...int) int {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

// Returns minimun value of a list of integers
func Min(values ...int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}