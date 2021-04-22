package array

// ReverseString reverses the Unicode points in the data string
// to return a lexically reversed version of the string. It has time
// complexity O(n) and space complexity O(n).
func ReverseString(data string) string {
	var reversedData string
	dataRunes := []rune(data)

	for i := len(dataRunes) - 1; i > -1; i-- {
		reversedData += string(dataRunes[i])
	}

	return reversedData
}
