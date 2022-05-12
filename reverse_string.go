package reverse_string

func ReverseString(input string) (output string) {
	for i := range input {
		output = output + string(input[len(input)-1-i])
	}
	return output
}
