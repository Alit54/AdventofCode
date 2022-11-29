package util

func removeLastElement(slice []string) []string {
	return append(slice[:len(slice)-1])
}
