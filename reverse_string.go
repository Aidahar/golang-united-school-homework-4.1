package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	strLen := len(input) - 1
	bytStr := []byte(input)
	var ans []byte
	for i := strLen; i >= 0; i-- {
		ans = append(ans, bytStr[i])
	}
	output = string(ans)
	return output
}
