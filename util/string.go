package util

import (
	"sort"
	"strings"
)

// Returns a strings ordered in alfabetical order.
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// Returns a slice of strings divided in n parts
/*func SplitEqually(s string, n int) []string {
	slice := make([]string, 0)
	for i := 0; i < n; i++ {
		slice = append(slice, s)
	}
	return slice
}*/

// Returns a slice of strings divided in the index position. Value in the index position will be in the second part.
func SplitStrings(s string, index int) []string {
	slice := make([]string, 0)
	slice = append(slice, s[:index])
	slice = append(slice, s[index:])
	return slice
}
