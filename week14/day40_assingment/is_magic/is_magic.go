package is_magic

func IsMagic(A int) int {
	if A == 1 {
		return 1
	} else if A < 10 {
		return 0
	}
	return IsMagic(DigitSum(A))

}
func DigitSum(A int) int {
	if A == 0 {
		return 0
	}
	return DigitSum(A/10) + A%10
}
