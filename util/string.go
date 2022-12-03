package util

import (
	"math"
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

// Returns a slice of runes with the common elements between 2 strings
func CommonRunes(s string, t string) []rune {
	slice := make([]rune, 0)
	// this for needs a fix using Alit.Min instead of math.Min
	for i := 0; i < int(math.Min(float64(len(s)), float64(len(t)))); i++ {
		if s[i] == t[i] {
			slice = append(slice, rune(s[i]))
		}
	}
	return slice
}

// Returns if a slice is all in uppercase
func IsUpper(s string) bool {
	return s == strings.ToUpper(s)
}

// Returns if a slice is all in uppercase
func IsLower(s string) bool {
	return s == strings.ToLower(s)
}
