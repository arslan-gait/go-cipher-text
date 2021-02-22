package main

var lowSlice = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var capitalSlice = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func getCharacterIndex(slice []rune, target rune) int {
	for iter, item := range slice {
		if target == item {
			return iter
		}
	}
	return -1
}

func shiftRune(curIdx int, iter int, result []rune, key int, alphabet []rune) {

	newIdx := (curIdx + key) % len(alphabet)

	if newIdx < 0 {
		newIdx += len(alphabet)
	}

	result[iter] = alphabet[newIdx]
}

func encryptDecryptText(text *string, key int, firstAlphabet []rune, secondAlphabet []rune, operation string) string {

	if operation == "decrypt" {
		key *= -1
	}

	result := []rune(*text)

	for iter, item := range result {
		if curIdx := getCharacterIndex(firstAlphabet, item); curIdx != -1 {
			shiftRune(curIdx, iter, result, key, firstAlphabet)
		} else if curIdx := getCharacterIndex(secondAlphabet, item); curIdx != -1 {
			shiftRune(curIdx, iter, result, key, secondAlphabet)
		}
	}
	return string(result)
}
