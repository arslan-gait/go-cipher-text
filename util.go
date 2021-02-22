package main

var lowSlice = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var capitalSlice = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func getCharacterIndex(slice []rune, target rune) int {
	for index, char := range slice {
		if target == char {
			return index
		}
	}
	return -1
}

func encryptDecryptText(text *string, key int, slice1 []rune, slice2 []rune, operation rune) string {
	if operation == 'd' {
		key *= -1
	}

	runes := []rune(*text)

	for index, char := range runes {
		lowSliceSize, capitalSliceSize := len(slice1), len(slice2)

		if curIdx := getCharacterIndex(slice1, char); curIdx != -1 {
			newIdx := (curIdx + key) % lowSliceSize
			if newIdx < 0 {
				newIdx += lowSliceSize
			}
			runes[index] = slice1[newIdx]

		} else if curIdx := getCharacterIndex(slice2, char); curIdx != -1 {
			newIdx := (curIdx + key) % capitalSliceSize
			if newIdx < 0 {
				newIdx += capitalSliceSize
			}
			runes[index] = slice2[newIdx]
		}
	}
	return string(runes)
}
