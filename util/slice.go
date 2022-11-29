package util

func RemoveLastElement(slice []any) []any {
	return append(slice[:len(slice)-1])
}
