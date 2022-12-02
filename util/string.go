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
