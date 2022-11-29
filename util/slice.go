package util

func RemoveLastElement(slice []string) []string {
	return append(slice[:len(slice)-1])
}
