package leap_year

func FindLeapYear(A int) int {
	if A%400 == 0 || (A%4 == 0 && A%100 != 0) {
		return 1
	}
	return 0
}
