package array

// ReverseString reverses the Unicode points in the data string
// to return a lexically reversed version of the string. It has time
// complexity O(n) and space complexity O(n).
func ReverseString(data string) string {
	var reversedData string
	dataRunes := []rune(data)
	dataLen := len(dataRunes)

	for i := range dataRunes {
		reversedData += string(dataRunes[dataLen-i-1])
	}

	return reversedData
}
