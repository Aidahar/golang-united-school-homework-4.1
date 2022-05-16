package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	runStr := []rune(input)
	strLen := len(runStr) - 1
	var ans []rune
	for i := strLen; i >= 0; i-- {
		ans = append(ans, runStr[i])
	}
	output = string(ans)
	return output
}
