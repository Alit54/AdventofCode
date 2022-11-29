package util

// Returns the slice of the input without the last element.
func RemoveLastElement(slice []any) []any {
	return append(slice[:len(slice)-1])
}

// Returns the slice of the input without the first element.
func RemoveFirstElement(slice []any) []any {
	return append(slice[1:])
}

// Returns the slice of the input without the indicated element. Panics if index is less than 0.
func RemoveElement(slice []any, index int) []any {
	if index < 0 {
		panic("Index can't be less than 0")
	}
	support := slice[:index]
	return append(support, slice[index+1:]...)
}
