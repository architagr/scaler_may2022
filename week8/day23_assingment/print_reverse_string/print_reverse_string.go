package print_reverse_string

func ReverseString(str string, n int) string {

	if n == 0 {
		return string(str[n])
	}
	return string(str[n]) + ReverseString(str, n-1)
}
