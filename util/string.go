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

// Returns a slice of strings divided in n parts.
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

// Returns a slice of runes with the common elements between 2 strings.
func CommonRunes2(s string, t string) []rune {
	slice := make([]rune, 0)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] && !CheckRunes(slice, rune(s[i])) {
				slice = append(slice, rune(s[i]))
			}
		}
	}
	return slice
}

// Returns a slice of runes with the common elements between 3 strings.
func CommonRunes3(s string, t string, u string) []rune {
	slice := make([]rune, 0)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			for k := 0; k < len(u); k++ {
				if s[i] == t[j] && t[j] == u[k] && !CheckRunes(slice, rune(s[i])) {
					slice = append(slice, rune(s[i]))
				}
			}
		}
	}
	return slice
}

// Returns TRUE if r is inside slice, FALSE otherwise.
func CheckRunes(slice []rune, r rune) bool {
	for _, v := range slice {
		if v == r {
			return true
		}
	}
	return false
}

// Returns if a slice is all in uppercase
func IsUpper(s string) bool {
	return s == strings.ToUpper(s)
}

// Returns if a slice is all in uppercase
func IsLower(s string) bool {
	return s == strings.ToLower(s)
}
